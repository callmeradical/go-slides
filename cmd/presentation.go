// Copyright © 2017 Lars Cromley lars@callmeradical.com
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
	"runtime/debug"

	"github.com/callmeradical/go-slides/lib"
	"github.com/spf13/cobra"
)

var presentationName string
var presentationCache bool

// presentationCmd represents the presentation command
var presentationCmd = &cobra.Command{
	Use:   "presentation",
	Short: "Create new presentation",
	Long:  `Create a new presentation. Creates a directory, downloads webslides and unpacks it.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		err := lib.Presentation(version, presentationName, presentationCache)
		if err != nil {
			debug.PrintStack()
			log.Println(err.Error())
		}

		fmt.Println(fmt.Sprintf("Created presentation %s", presentationName))

	},
}

func init() {
	newCmd.AddCommand(presentationCmd)

	presentationCmd.Flags().StringVarP(&version, "version", "v", "1.3.1", "specify the version of webslides to download")
	presentationCmd.Flags().StringVarP(&presentationName, "name", "n", "", "name of new presentation")
	presentationCmd.Flags().
		presentationCmd.Flags().BoolP("cache", "c", true, "specify whether or not to cache the version of webslides")

}
