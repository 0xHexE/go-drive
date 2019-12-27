/*
 * Copyright (c) 2019 Omkar Yadav. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package storage_client

import (
	"context"
	"io"
)

type StorageClient interface {
	PutObject(ctx context.Context, fileName string, reader io.Reader, size int64) (error, int64)
	GetObject(ctx context.Context, objectName string) (error, io.Reader)
}
