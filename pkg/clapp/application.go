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

// func (a *Application) initializeConfig(cmd *cobra.Command, args []string) error {
// 	return a.bindFlags(cmd)
// }

// func (a *Application) bindFlags(cmd *cobra.Command) error {
// 	var err error
// 	cmd.Flags().VisitAll(func(f *pflag.Flag) {
// 		// Determine the naming convention of the flags when represented in the config file
// 		configName := f.Name
//
// 		// If using camelCase in the config file, replace hyphens with a camelCased string.
// 		// Since viper does case-insensitive comparisons, we don't need to bother fixing the case, and only need to remove the hyphens.
// 		if a.config.ReplaceHyphenWithCamelCase() {
// 			configName = strings.ReplaceAll(f.Name, "-", "")
// 		}
//
// 		// Apply the viper config value to the flag when the flag is not set and viper has a value
// 		if !f.Changed && a.config.Viper().IsSet(configName) {
// 			val := a.config.Viper().Get(configName)
// 			err = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
// 			if err != nil {
// 				slog.Error("could not update flag from viper", "error", err)
// 			}
// 		}
// 	})
// 	return err
// }

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

	rootCmd.PersistentFlags().StringVarP(&a.configFileFlag, "file", "f", "config.yaml", "config filename")
	rootCmd.PersistentFlags().StringVarP(&a.configPathFlag, "path", "p", "", "config file path")

	a.Command = rootCmd
	return a
}

//
// func NewApplicationWithConfiguration(use string, short string, long string, version string, config *Configuration) *Application {
// 	a := &Application{
// 		// config: config,
// 	}
//
// 	rootCmd := &cobra.Command{
// 		Use:   use,
// 		Short: short,
// 		Long:  long,
// 		// PersistentPreRunE: a.initializeConfig,
// 		Args:    cobra.MinimumNArgs(1),
// 		Version: version,
// 	}
//
// 	rootCmd.Flags().StringVarP(&a.configFileFlag, "file", "f", "config.yaml", "config filename")
// 	rootCmd.Flags().StringVarP(&a.configPathFlag, "path", "p", "", "config file path")
//
// 	a.Command = rootCmd
// 	return a
// }
//
// func NewApplicationWithCommand(command *cobra.Command, config *Configuration) *Application {
// 	return &Application{
// 		Command: command,
// 		// config:  config,
// 	}
// }
