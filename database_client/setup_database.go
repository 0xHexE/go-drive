/*
 * Copyright (c) 2019 Omkar Yadav. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package database_client

import (
	"github.com/httpsOmkar/go-drive/env_config"
	"github.com/jinzhu/gorm"
)

type Database struct {
	Database *gorm.DB
}

type File struct {
	gorm.Model
	FileId   string `sql:"index"`
	FileName string
}

// Init the database
func InitDatabase(config *env_config.AppEnvConfig) (error, *Database) {
	db, err := gorm.Open(config.Database.Type, config.Database.ConnectionUrl)
	if err != nil {
		return err, nil
	}

	database := Database{
		Database: db,
	}

	database.setup()

	return nil, &database
}

// setup the database
func (database *Database) setup() {}
