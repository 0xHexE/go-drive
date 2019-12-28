/*
 * Copyright (c) 2019 Omkar Yadav. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package env_config

import "github.com/spf13/viper"

type AppEnvConfig struct {
	Server   ServerConfig
	Database DatabaseConnection
	Minio    MinioConfig
}

type MinioConfig struct {
	Endpoint, AccessKey, SecretAccessKey, Location, Bucket string
	UseSSL                                                 bool
}

type ServerConfig struct {
	Port int
}

type DatabaseConnection struct {
	Type          string
	ConnectionUrl string
}

func LoadConfig() (error, *AppEnvConfig) {
	// Load config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err, nil
	}

	return nil, nil
}
