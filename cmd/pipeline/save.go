// Copyright (c) 2018, Google, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package pipeline

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/spinnaker/spin/util"
)

type saveOptions struct {
	*PipelineOptions
	pipelineFile string
}

var (
	savePipelineShort = "Save the provided pipeline"
	savePipelineLong  = "Save the provided pipeline"
)

func NewSaveCmd(pipelineOptions *PipelineOptions) *cobra.Command {
	options := &saveOptions{
		PipelineOptions: pipelineOptions,
	}
	cmd := &cobra.Command{
		Use:     "save",
		Aliases: []string{},
		Short:   savePipelineShort,
		Long:    savePipelineLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			return savePipeline(options)
		},
	}

	cmd.PersistentFlags().StringVarP(&options.pipelineFile, "file", "f", "", "path to the pipeline file")

	return cmd
}

func savePipeline(options *saveOptions) error {
	pipelineJSON, err := util.ParseJSONFromFileOrStdin(options.pipelineFile, false)
	if err != nil {
		return err
	}
	valid := true
	if _, exists := pipelineJSON["name"]; !exists {
		options.UI.Error("Required pipeline key 'name' missing...\n")
		valid = false
	}

	if _, exists := pipelineJSON["application"]; !exists {
		options.UI.Error("Required pipeline key 'application' missing...\n")
		valid = false
	}

	if template, exists := pipelineJSON["template"]; exists && len(template.(map[string]interface{})) > 0 {
		if _, exists := pipelineJSON["schema"]; !exists {
			options.UI.Error("Required pipeline key 'schema' missing for templated pipeline...\n")
			valid = false
		}
		pipelineJSON["type"] = "templatedPipeline"
	}

	if !valid {
		return fmt.Errorf("Submitted pipeline is invalid: %s\n", pipelineJSON)
	}
	application := pipelineJSON["application"].(string)
	pipelineName := pipelineJSON["name"].(string)

	foundPipeline, queryResp, _ := options.GateClient.ApplicationControllerApi.GetPipelineConfigUsingGET(options.GateClient.Context, application, pipelineName)

	if queryResp.StatusCode != http.StatusOK {
		return fmt.Errorf("Encountered an error querying pipeline, status code: %d\n", queryResp.StatusCode)
	}

	_, exists := pipelineJSON["id"].(string)
	var foundPipelineID string
	if len(foundPipeline) > 0 {
		foundPipelineID = foundPipeline["id"].(string)
	}
	if !exists && foundPipelineID != "" {
		pipelineJSON["id"] = foundPipelineID
	}

	saveResp, saveErr := options.GateClient.PipelineControllerApi.SavePipelineUsingPOST(options.GateClient.Context, pipelineJSON)

	if saveErr != nil {
		return saveErr
	}
	if saveResp.StatusCode != http.StatusOK {
		return fmt.Errorf("Encountered an error saving pipeline, status code: %d\n", saveResp.StatusCode)
	}

	options.UI.Success("Pipeline save succeeded")
	return nil
}
