// Copyright © 2017 Luiz Filho
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/luizbafilho/gosubst/gosubst"
	"github.com/spf13/cobra"
)

var (
	valuesFile string
	valuesType string
	tplPaths   []string
	outPath    string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gosubst",
	Short: "substitutes values of variables file in template files",
	Long:  `gosubst copies stardard input to standard output replacing all variables present in values file`,
	Run: func(cmd *cobra.Command, args []string) {
		subst, err := gosubst.NewSubst(valuesFile, valuesType, os.Stdin, os.Stdout)
		if err != nil {
			log.Fatal(err)
		}

		if err := subst.Render(); err != nil {
			log.Fatal(err)
		}
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&valuesFile, "values", "v", "", "values file")
	RootCmd.PersistentFlags().StringVar(&valuesType, "type", "yaml", "values type (toml, yaml or json)")
}
