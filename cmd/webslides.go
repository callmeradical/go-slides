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
	"log"
	"os"

	"github.com/callmeradical/go-slides/lib"
	"github.com/spf13/cobra"
)

var version string

// webslidesCmd represents the webslides command
var webslidesCmd = &cobra.Command{
	Use:   "webslides",
	Short: "fetch webslides from the internet",
	Long: `Webslides is a great project with lots of builtin classes
	and code snipets, this command grabs the specified version of webslides`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		wd, err := os.Getwd()
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}

		err = lib.FetchSlides(version, wd, true)
		if err != nil {
			log.Println(err.Error())
		}
	},
}

func init() {
	getCmd.AddCommand(webslidesCmd)

	webslidesCmd.Flags().StringVarP(&version, "version", "v", "1.3.1", "specify the version of webslides to download")
}
