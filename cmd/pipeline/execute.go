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
	"errors"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spinnaker/spin/cmd/gateclient"
	"github.com/spinnaker/spin/util"
)

type ExecuteOptions struct {
	*pipelineOptions
	output        string
	application   string
	name          string
	parameterFile string
	artifactsFile string
}

var (
	executePipelineShort = "Execute the provided pipeline"
	executePipelineLong  = "Execute the provided pipeline"
)

func NewExecuteCmd(pipelineOptions pipelineOptions) *cobra.Command {
	options := ExecuteOptions{
		pipelineOptions: &pipelineOptions,
	}
	cmd := &cobra.Command{
		Use:     "execute",
		Aliases: []string{"exec"},
		Short:   executePipelineShort,
		Long:    executePipelineLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			return executePipeline(cmd, options)
		},
	}

	cmd.PersistentFlags().StringVarP(&options.application, "application", "a", "", "Spinnaker application the pipeline lives in")
	cmd.PersistentFlags().StringVarP(&options.name, "name", "n", "", "name of the pipeline to execute")
	cmd.PersistentFlags().StringVarP(&options.parameterFile, "parameter-file", "f", "", "file to load pipeline parameter values from")
	cmd.PersistentFlags().StringVarP(&options.artifactsFile, "artifacts-file", "t", "", "file to load pipeline artifacts from")

	return cmd
}

func executePipeline(cmd *cobra.Command, options ExecuteOptions) error {
	gateClient, err := gateclient.NewGateClient(cmd.InheritedFlags())
	if err != nil {
		return err
	}

	if options.application == "" || options.name == "" {
		return errors.New("one of required parameters 'application' or 'name' not set")
	}

	parameters := map[string]interface{}{}
	parameters, err = util.ParseJsonFromFile(options.parameterFile, true)
	if err != nil {
		return fmt.Errorf("Could not parse supplied pipeline parameters: %v.\n", err)
	}

	artifactsFile := map[string]interface{}{}
	artifactsFile, err = util.ParseJsonFromFile(options.artifactsFile, true)
	if err != nil {
		return fmt.Errorf("Could not parse supplied artifacts: %v.\n", err)
	}

	trigger := map[string]interface{}{"type": "manual"}
	if len(parameters) > 0 {
		trigger["parameters"] = parameters
	}

	if _, ok := artifactsFile["artifacts"]; ok {
		artifacts := artifactsFile["artifacts"].([]interface{})
		if len(artifacts) > 0 {
			trigger["artifacts"] = artifacts
		}
	}

	_, resp, err := gateClient.PipelineControllerApi.InvokePipelineConfigUsingPOST1(gateClient.Context,
		options.application,
		options.name,
		map[string]interface{}{"trigger": trigger})

	if err != nil {
		return fmt.Errorf("Execute pipeline failed with response: %v and error: %s\n", resp, err)
	}

	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Encountered an error executing pipeline, status code: %d\n", resp.StatusCode)
	}

	util.UI.Info(util.Colorize().Color(fmt.Sprintf("[reset][bold][green]Pipeline execution started")))

	return nil
}
