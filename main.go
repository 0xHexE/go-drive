/*
 * Copyright (c) 2019 Omkar Yadav. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"github.com/httpsOmkar/go-drive/app"
	"github.com/httpsOmkar/go-drive/env_config"
	"github.com/httpsOmkar/go-drive/http_server"
	"github.com/httpsOmkar/go-drive/storage_client"
	"log"
)

func main() {
	err, appConfig := env_config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to read AppConfig %v", err)
	}

	log.Println("Setting up S3 client")
	err, _ = storage_client.InitMinio(appConfig)

	if err != nil {
		log.Fatalf("Failed to setup minio client %v", err)
	}

	log.Println("S3 server connected")

	appInstance := app.App{
		AppEnvConfig: appConfig,
	}

	http_server.InitHttp(&appInstance)
}
