// Copyright 2019 FairwindsOps Inc
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
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"k8s.io/klog"

	"github.com/fairwindsops/goldilocks/pkg/summary"
	"github.com/fairwindsops/goldilocks/pkg/utils"
)

var excludeContainers string
var outputFile string

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.PersistentFlags().StringVarP(&excludeContainers, "exclude-containers", "e", "", "Comma delimited list of containers to exclude from recommendations.")
	summaryCmd.PersistentFlags().StringVarP(&outputFile, "output-file", "f", "", "File to write output from audit.")
}

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Genarate a summary of the vpa recommendations in a namespace.",
	Long:  `Gather all the vpa data in a namespace and generaate a summary of the recommendations.`,
	Run: func(cmd *cobra.Command, args []string) {

		data, _ := summary.GetInstance().Run(utils.VpaLabels, excludeContainers)
		summaryJSON, err := json.Marshal(data)
		if err != nil {
			klog.Fatalf("Error marshalling JSON: %v", err)
		}

		if outputFile != "" {
			err := ioutil.WriteFile(outputFile, summaryJSON, 0644)
			if err != nil {
				klog.Fatalf("Failed to write summary to file: %v", err)
			}

			fmt.Println("Summary has been written to", outputFile)

		} else {
			fmt.Println(string(summaryJSON))
		}
	},
}
