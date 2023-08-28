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
	"os"

	"github.com/spf13/cobra"
)

type Application struct {
	*cobra.Command
	// config Configuration

	configFileFlag string
	configPathFlag string
}

func (a *Application) RegisterCommands(c []Commander) {
	for _, cmdr := range c {
		a.AddCommand(cmdr.Initialize())
	}
}

func (a *Application) Run() {
	if err := a.Execute(); err != nil {
		slog.Error("could not start application", "error", err)
		os.Exit(1)
	}
}

func NewApplication(use string, short string, long string, version string) *Application {
	a := &Application{
		// config: nil,
	}

	rootCmd := &cobra.Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Args:    cobra.MinimumNArgs(1),
		Version: version,
	}

	rootCmd.PersistentFlags().StringVarP(&a.configFileFlag, "file", "f", "config.yaml", "config file name")
	rootCmd.PersistentFlags().StringVarP(&a.configPathFlag, "path", "p", "", "config file path")

	a.Command = rootCmd
	return a
}
