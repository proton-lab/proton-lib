// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/proton-lab/autom/linuxAP/app/cmdclient"
	"github.com/proton-lab/autom/linuxAP/app/common"
	"github.com/spf13/cobra"
	"log"
)

// showCmd represents the show command
var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "show " + ProgramName + " config",
	Long:  "show " + ProgramName + " config",
	Run: func(cmd *cobra.Command, args []string) {

		if remoteaddr == "" {
			if _, err := common.IsLinuxAPProcessStarted(); err != nil {
				log.Println(err)
				return
			}
		}
		cmdclient.DefaultCmdSend(remoteaddr, common.CMD_CONFIG_SHOW)
	},
}

func init() {
	configCmd.AddCommand(configShowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
