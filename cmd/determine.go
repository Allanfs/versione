/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/allanfs/versione/internal/app"
)

// determineCmd represents the determine command
var determineCmd = &cobra.Command{
	Use:   "determine",
	Short: "determina a proxima versão",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		VERSAO_ATUAL, _ = cmd.Flags().GetString("versao")
		BRANCH_ORIGEM, _ = cmd.Flags().GetString("branch")

		regexes := []string{}
		regexes = append(regexes, viper.GetStringSlice("major.patterns")...)
		regexes = append(regexes, viper.GetStringSlice("minor.patterns")...)
		regexes = append(regexes, viper.GetStringSlice("patch.patterns")...)

		// viper.GetStringMap("versions")["minor"].(map[string]interface{})["patterns"]
		var novaVersao string
		for k := range viper.GetStringMap("versions") {

			if app.StringRegexMatchBranch(viper.GetStringMapStringSlice("versions." + k)["patterns"], BRANCH_ORIGEM) {
				novaVersao = app.GetUpdaterByName(k).Update(VERSAO_ATUAL)
				break
			}
		}

		fmt.Println(novaVersao)

	},
}

var (
	VERSAO_ATUAL  string
	BRANCH_ORIGEM string
)

func init() {
	rootCmd.AddCommand(determineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// determineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	determineCmd.Flags().StringP("versao", "v", "0.1.0", "versão atual do projeto")
	determineCmd.Flags().StringP("branch", "b", "", "branch de origem")

	determineCmd.MarkFlagRequired("versao")
	determineCmd.MarkFlagRequired("branch")

}
