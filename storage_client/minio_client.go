/*
 * Copyright (c) 2019 Omkar Yadav. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package storage_client

import (
	"context"
	"github.com/httpsOmkar/go-drive/env_config"
	"github.com/minio/minio-go/v6"
	"io"
)

type MinioClient struct {
	Client *minio.Client
	config *env_config.AppEnvConfig
}

// Put object into the bucket
func (client *MinioClient) PutObject(ctx context.Context, fileName string, reader io.Reader, size int64) (error, int64) {
	count, err := client.Client.PutObjectWithContext(
		ctx,
		client.config.Minio.Bucket,
		fileName, reader, size,
		minio.PutObjectOptions{},
	)
	return err, count
}

// Get object from the client
func (client *MinioClient) GetObject(ctx context.Context, objectName string) (error, io.Reader) {
	object, err := client.Client.GetObjectWithContext(ctx, client.config.Minio.Bucket, objectName, minio.GetObjectOptions{})

	var returnValue io.Reader
	returnValue = object

	return err, returnValue
}

// Setup the S3 bucket for the initial use
// make bucket for the first use. Create
// bucket if not exists
func (client *MinioClient) setup() error {
	config := client.config
	bucket := client.config.Minio.Bucket

	exists, err := client.Client.BucketExists(bucket)

	if err != nil {
		return err
	}

	if !exists {
		err = client.Client.MakeBucket(config.Minio.Bucket, config.Minio.Location)

		if err != nil {
			return err
		}
	}

	return nil
}

// Setup minio client and setup basic
// setup on the minio
func InitMinio(config *env_config.AppEnvConfig) (error, *MinioClient) {
	minioClient := MinioClient{}

	// Create minio client
	client, err := minio.New(
		config.Minio.Endpoint,
		config.Minio.AccessKey,
		config.Minio.SecretAccessKey,
		config.Minio.UseSSL,
	)

	if err != nil {
		return err, nil
	}

	minioClient.config = config
	minioClient.Client = client

	err = minioClient.setup()

	if err != nil {
		return err, nil
	}

	return nil, &minioClient
}
