// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

	"github.com/spf13/cobra"
)

// laterCmd represents the later command
var laterCmd = &cobra.Command{
	Use:   "later",
	Short: "Move a task back to TODO",
	Long: `If you change your mind about doing a task today, 
or if you couldn't do it as planned, you can always move 
the task back to TODO using the later command`,
	Run: func(cmd *cobra.Command, args []string) {
		id = args[0]
		fmt.Println("later called on", id)
	},
}

func init() {
	RootCmd.AddCommand(laterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// laterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// laterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
