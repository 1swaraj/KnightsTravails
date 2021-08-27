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
	"bytes"
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// setSizeCmd represents the setSize command
var setSizeCmd = &cobra.Command{
	Use:   "set-size",
	Short: "Set board size",
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigType("yaml")
		var config = []byte(`
board_rows: `+args[0]+`
board_cols: `+args[1]+`
`)
		viper.ReadConfig(bytes.NewBuffer(config))
		viper.WriteConfigAs("./config")

		err := viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setSizeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setSizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setSizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
