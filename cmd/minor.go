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
	"gitlab.com/allanfs/versione/internal/app"
)

// minorCmd represents the minor command
var minorCmd = &cobra.Command{
	Use:   "minor",
	Short: "Incrementa a vesão do meio: #.X.#",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		updater := app.MinorUpdater{}
		versao, _ := cmd.Flags().GetString("versao")
		fmt.Println(updater.Update(versao))
	},
}

func init() {
	bumpCmd.AddCommand(minorCmd)
	minorCmd.Flags().StringP("versao", "v", "0.1.0", "versão atual do projeto")

	minorCmd.MarkFlagRequired("versao")
}
