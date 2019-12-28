/*
 * Copyright (c) 2019 Omkar Yadav. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"github.com/httpsOmkar/go-drive/app"
	"github.com/httpsOmkar/go-drive/database_client"
	"github.com/httpsOmkar/go-drive/env_config"
	"github.com/httpsOmkar/go-drive/http_client"
	"github.com/httpsOmkar/go-drive/storage_client"
	"log"
)

func main() {
	err, appConfig := env_config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to read AppConfig %v", err)
	}

	log.Println("Setting up S3 client")
	err, storageClient := storage_client.InitMinio(appConfig)

	if err != nil {
		log.Fatalf("Failed to setup minio client %v", err)
	}

	log.Println("S3 server connected")

	log.Println("Connecting to SQL Database")

	err, databaseClient := database_client.InitDatabase(appConfig)

	if err != nil {
		log.Fatalf("Failed to connect to sql %v", err)
	}

	app := app.App{
		Database:     databaseClient,
		Storage:      storageClient,
		AppEnvConfig: appConfig,
	}

	http_client.InitHttp(app)
}
