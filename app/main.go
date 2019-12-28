/*
 * Copyright (c) 2019 Omkar Yadav. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package app

import (
	"context"
	"github.com/google/uuid"
	"github.com/httpsOmkar/go-drive/database_client"
	"github.com/httpsOmkar/go-drive/env_config"
	"github.com/httpsOmkar/go-drive/storage_client"
	"github.com/jinzhu/gorm"
	"io"
)

type App struct {
	Database     *database_client.Database
	Storage      *storage_client.StorageClient
	AppEnvConfig *env_config.AppEnvConfig
}

// Basic file upload into the database and S3
func (app *App) UploadFile(ctx context.Context, reader io.Reader, size int64) (error, *string) {
	fileId := uuid.New().String()
	err, _ := (*app.Storage).PutObject(ctx, fileId, reader, size)
	if err != nil {
		return err, nil
	}

	fileRecord := database_client.File{
		Model:    gorm.Model{},
		FileId:   "",
		FileName: "",
	}

	app.Database.Database.Create(&fileRecord)

	return nil, &fileId
}

func (app *App) DownloadFile() (error, *string) {
	return nil, nil
}
