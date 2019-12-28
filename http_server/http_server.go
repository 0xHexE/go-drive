/*
 * Copyright (c) 2019 Omkar Yadav. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package http_server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/httpsOmkar/go-drive/app"
	"log"
	"net/http"
	"time"
)

func InitHttp(config *app.App) {
	router := mux.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr: fmt.Sprintf(
			"%s:%d",
			config.AppEnvConfig.Server.ListenAddress,
			config.AppEnvConfig.Server.Port,
		),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
