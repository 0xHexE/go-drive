/*
 * Copyright (c) 2019 Omkar Yadav. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"github.com/httpsOmkar/go-drive/app"
	"github.com/httpsOmkar/go-drive/env_config"
	"github.com/httpsOmkar/go-drive/http_server"
	"github.com/httpsOmkar/go-drive/storage_client"
	"log"
	"net/http"
	"time"
)

func main() {
	err, appConfig := env_config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to read AppConfig %v", err)
	}

	log.Println("Setting up S3 client")
	err, storage := storage_client.InitMinio(appConfig)

	if err != nil {
		log.Fatalf("Failed to setup minio client %v", err)
	}

	log.Println("S3 server connected")

	var storageClient storage_client.StorageClient

	storageClient = storage

	appInstance := app.App{
		AppEnvConfig: appConfig,
		Storage:      &storageClient,
	}

	server := http_server.InitHttp(&appInstance)

	srv := &http.Server{
		Handler: server,
		Addr: fmt.Sprintf(
			"%s:%d",
			appConfig.Server.ListenAddress,
			appConfig.Server.Port,
		),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
