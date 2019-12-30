/*
 * Copyright (c) 2019 Omkar Yadav. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package env_config

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppEnvConfig struct {
	Server   ServerConfig
	Minio    MinioConfig
	Endpoint Endpoint
	Upload   UploadConfig
}

type MinioConfig struct {
	Endpoint, AccessKey, SecretAccessKey, Location, Bucket string
	UseSSL                                                 bool
}

type ServerConfig struct {
	Port          int
	ListenAddress string
}

type UploadConfig struct {
	MaxUploadSize int64
	FilePath      string
}

type Endpoint struct {
	ApiVersion, UploadUrl, DownloadUrl string
}

func (endpoint *Endpoint) GenUrl(endPoint string) string {
	return fmt.Sprintf("/api/%s/%s", endpoint.ApiVersion, endPoint)
}

func LoadConfig() (error, *AppEnvConfig) {
	// Load config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error
		} else {
			return err, nil
		}
	}

	var appEnvConfig AppEnvConfig
	err := viper.Unmarshal(&appEnvConfig)

	if err != nil {
		return err, nil
	}

	return nil, &appEnvConfig
}
