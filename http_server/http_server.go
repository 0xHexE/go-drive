/*
 * Copyright (c) 2019 Omkar Yadav. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package http_server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/httpsOmkar/go-drive/app"
	"io"
	"net/http"
	"path"
)

type HandlerFunc func(writer http.ResponseWriter, request *http.Request)

func InitHttp(config *app.App) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(
		config.AppEnvConfig.Endpoint.GenUrl(config.AppEnvConfig.Endpoint.DownloadUrl),
		HandleDownload(config),
	)

	fmt.Println(config.AppEnvConfig.Endpoint.GenUrl(config.AppEnvConfig.Endpoint.UploadUrl))

	router.HandleFunc(
		config.AppEnvConfig.Endpoint.GenUrl(config.AppEnvConfig.Endpoint.UploadUrl),
		HandleUpload(config),
	)

	return router
}

func HandleDownload(app2 *app.App) HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		err, read := app2.DownloadFile(request.Context(), request.URL.Query().Get("file"))

		if err != nil {
			_, _ = writeJSON(writer, "{ \"error\": \"INTERNAL SERVER ERROR\" }", http.StatusInternalServerError)
			return
		}

		_, _ = io.Copy(writer, read)
	}
}

func HandleUpload(app *app.App) HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		formPath := request.URL.Query().Get("path")

		if formPath == "" {
			_, _ = writeJSON(writer, "{ \"error\": \"MISSING PATH IN QUERY\" }", http.StatusBadRequest)
			return
		}

		request.Body = http.MaxBytesReader(writer, request.Body, app.AppEnvConfig.Upload.MaxUploadSize)

		if err := request.ParseMultipartForm(app.AppEnvConfig.Upload.MaxUploadSize); err != nil {
			_, _ = writeJSON(writer, "{\n  \"error\": \"FILE TO LONG\"\n}", http.StatusBadRequest)
			return
		}

		file, info, err := request.FormFile(app.AppEnvConfig.Upload.FilePath)

		if err != nil {
			_, _ = writeJSON(writer, "{ \"error\": \"INVALID FILE\" }", http.StatusBadRequest)
			return
		}

		filePath := path.Join(formPath, info.Filename)

		defer func() {
			_ = file.Close()
		}()

		err = app.UploadFile(request.Context(), filePath, file, info.Size)

		if err != nil {
			_, _ = writeJSON(writer, "{ \"error\": \"INTERNAL SERVER ERROR\" }", http.StatusInternalServerError)
			return
		}

		returnValue := map[string]string{
			"filePath": filePath,
		}

		bytes, err := json.Marshal(returnValue)

		if err != nil {
			_, _ = writeJSON(writer, "{ \"error\": \"INTERNAL SERVER ERROR\" }", http.StatusInternalServerError)
			return
		}

		_, _ = writeJSONFromByte(writer, bytes, http.StatusCreated)
	}
}

func writeJSON(w http.ResponseWriter, message string, statusCode int) (int, error) {
	return writeJSONFromByte(w, []byte(message), statusCode)
}

func writeJSONFromByte(w http.ResponseWriter, message []byte, statusCode int) (int, error) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	return w.Write(message)
}
