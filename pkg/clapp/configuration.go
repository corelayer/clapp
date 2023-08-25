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
	"fmt"
	"log/slog"
	"strings"

	"github.com/spf13/viper"
)

type Configuration struct {
	*viper.Viper
	filename                   string
	paths                      []string
	replaceHyphenWithCamelCase bool
}

func (c Configuration) initialize() error {
	c.initConfigDetails()
	c.initConfigPaths()
	err := c.ReadInConfig()
	if err != nil {
		slog.Debug("Error reading config from viper", "error", err)
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (c Configuration) initConfigDetails() {
	parts := strings.Split(c.filename, ".")
	if len(parts) < 2 {
		parts = append(parts, "yaml")
	}
	c.SetConfigName(parts[0])
	c.SetConfigType(parts[1])
}

func (c Configuration) initConfigPaths() {
	for _, path := range c.paths {
		slog.Debug("adding path to config", "path", path)
		c.AddConfigPath(path)
	}
}

func NewConfiguration(filename string, paths []string, replaceHyphenWithCamelCase bool) (*Configuration, error) {
	c := &Configuration{
		Viper:                      viper.New(),
		filename:                   filename,
		paths:                      paths,
		replaceHyphenWithCamelCase: replaceHyphenWithCamelCase,
	}

	err := c.initialize()
	if err != nil {
		slog.Error("could not initialize configuration", "error", err)
		return nil, fmt.Errorf("could not initialize configuration: %w", err)
	}
	return c, nil
}
