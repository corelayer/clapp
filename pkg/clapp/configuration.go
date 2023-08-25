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

// import (
// 	"github.com/spf13/viper"
// )
//
// type Configuration struct {
// 	*viper.Viper
// 	filename string
// 	path     string
// 	paths    []string
// 	// autoLoad                   bool
// 	// replaceHyphenWithCamelCase bool
// 	// initialized                bool
// 	// loaded                     bool
// }

//	func (c Configuration) FileName() *string {
//		return &c.filename
//	}
//
//	func (c Configuration) FilePath() *string {
//		return &c.path
//	}
//
//	func (c Configuration) Initialize() {
//		c.initConfigDetails()
//		c.initConfigPaths()
//
//		c.initialized = true
//	}
//
//	func (c Configuration) IsInitialized() bool {
//		return c.initialized
//	}
//
//	func (c Configuration) IsLoaded() bool {
//		return c.loaded
//	}
//
//	func (c Configuration) Load() error {
//		if !c.IsInitialized() {
//			c.Initialize()
//		}
//
//		err := c.viper.ReadInConfig()
//		if err != nil {
//			slog.Debug("error reading config", "error", err)
//			return fmt.Errorf("%w", err)
//		}
//		c.loaded = true
//		return nil
//	}
//
//	func (c Configuration) ReplaceHyphenWithCamelCase() bool {
//		return c.replaceHyphenWithCamelCase
//	}
//
//	func (c Configuration) Viper() *viper.Viper {
//		return c.viper
//	}
//
//	func (c Configuration) initConfigDetails() {
//		parts := strings.Split(c.filename, ".")
//		if len(parts) < 2 {
//			parts = append(parts, "yaml")
//		}
//		slog.Debug("setting config file name", "filename", parts[0])
//		c.viper.SetConfigName(parts[0])
//		slog.Debug("setting config file type", "filetype", parts[1])
//		c.viper.SetConfigType(parts[1])
//	}
//
//	func (c Configuration) initConfigPaths() {
//		if c.path != "" {
//			c.viper.AddConfigPath(c.path)
//		}
//
//		for _, path := range c.paths {
//			slog.Debug("adding config search path", "path", path)
//			c.viper.AddConfigPath(path)
//		}
//	}
// func NewConfiguration(filename string, path string, searchPaths []string, autoLoad bool, replaceHyphenWithCamelCase bool) (*Configuration, error) {
// 	var err error
// 	c := &Configuration{
// 		viper:    viper.New(),
// 		filename: filename,
// 		path:     path,
// 		paths:    searchPaths,
// 		autoLoad: autoLoad,
// 		// replaceHyphenWithCamelCase: replaceHyphenWithCamelCase,
// 		// initialized:                false,
// 		// loaded:                     false,
// 	}
//
// 	if autoLoad {
// 		err = c.Load()
// 	}
//
// 	return c, err
// }
