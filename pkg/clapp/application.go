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
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Application struct {
	*cobra.Command
	config *Configuration
}

func (a *Application) initializeConfig(cmd *cobra.Command, args []string) error {
	return a.bindFlags(cmd)
}

func (a *Application) bindFlags(cmd *cobra.Command) error {
	var err error
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Determine the naming convention of the flags when represented in the config file
		configName := f.Name

		// If using camelCase in the config file, replace hyphens with a camelCased string.
		// Since viper does case-insensitive comparisons, we don't need to bother fixing the case, and only need to remove the hyphens.
		if a.config.replaceHyphenWithCamelCase {
			configName = strings.ReplaceAll(f.Name, "-", "")
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && a.config.IsSet(configName) {
			val := a.config.Get(configName)
			err = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
			if err != nil {
				slog.Error("could not update flag from viper", "error", err)
			}
		}
	})
	return err
}

func (a *Application) RegisterCommands(c []Commander) {
	for _, cmdr := range c {
		a.AddCommand(cmdr.Initialize())
	}
}

func (a *Application) Run() {
	if err := a.Execute(); err != nil {
		slog.Error("Could not start application", "error", err)
		os.Exit(1)
	}
}

func NewApplication(use string, short string, long string, config *Configuration) *Application {
	a := &Application{
		config: config,
	}

	rootCmd := &cobra.Command{
		Use:               use,
		Short:             short,
		Long:              long,
		PersistentPreRunE: a.initializeConfig,
	}

	a.Command = rootCmd
	return a
}
