/*
 * Copyright (c) 2019 Omkar Yadav. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package auth

import "gopkg.in/yaml.v2"

type Config struct {
	ApiVersion, Kind string
	Metadata         ConfigMetadata
	Spec             ConfigSpec
}

type ConfigSpec struct {
	Match  string
	Access ConfigAccess
}

type ConfigAccess struct {
	Read, Write, Delete string
}

type ConfigMetadata struct {
	Name, Namespace string
}

// Parse config from yaml
func ParseConfig(input []byte) (error, *Config) {
	var config Config
	err := yaml.Unmarshal(input, &config)
	return err, &config
}
