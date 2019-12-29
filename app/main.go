/*
 * Copyright (c) 2019 Omkar Yadav. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package app

import (
	"context"
	"github.com/httpsOmkar/go-drive/env_config"
	"github.com/httpsOmkar/go-drive/storage_client"
	"io"
)

type App struct {
	Storage      *storage_client.StorageClient
	AppEnvConfig *env_config.AppEnvConfig
}

// Basic file upload into the database and S3
func (app *App) UploadFile(ctx context.Context, fileName string, reader io.Reader, size int64) error {
	err, _ := (*app.Storage).PutObject(ctx, fileName, reader, size)
	return err
}

func (app *App) DownloadFile() (error, *string) {
	return nil, nil
}
