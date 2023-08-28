/*
 * Copyright 2023 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package clapp

import (
	"log/slog"
	"strings"

	"github.com/spf13/viper"
)

type Configuration struct {
	filename string
	path     string
	paths    []string
	Viper    *viper.Viper
}

func (c Configuration) initConfigDetails() (string, string) {
	parts := strings.Split(c.filename, ".")
	if len(parts) < 2 {
		parts = append(parts, "yaml")
	}
	return parts[0], parts[1]
}

func (c Configuration) Initialize() {
	c.Viper = viper.New()

	if c.path != "" {
		fullPath := c.path + "/" + c.filename
		slog.Debug("setting config file", "file", fullPath)
		c.Viper.SetConfigFile(fullPath)
	} else {
		configName, configType := c.initConfigDetails()
		slog.Debug("setting config name", "name", configName)
		c.Viper.SetConfigName(configName)
		slog.Debug("setting config type", "type", configType)
		c.Viper.SetConfigType(configType)

		for _, path := range c.paths {
			slog.Debug("adding config search path", "path", path)
			c.Viper.AddConfigPath(path)
		}
	}
}

func (c Configuration) IsInitialized() bool {
	if c.Viper == nil {
		return false
	}
	return true
}

func NewConfiguration(filename string, path string, searchPaths []string, initialize bool) *Configuration {
	c := &Configuration{
		filename: filename,
		path:     path,
		paths:    searchPaths,
	}

	if initialize {
		c.Initialize()
	}

	return c
}
