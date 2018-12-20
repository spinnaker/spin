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
	"strings"
	"time"

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
}

var (
	executePipelineShort   = "Execute the provided pipeline"
	executePipelineLong    = "Execute the provided pipeline"
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
	parameters, err = util.ParseJsonFromFileOrStdin(options.parameterFile)
	if err != nil && strings.HasPrefix(err.Error(), "No json input") {
		// Pipeline can be executed with no parameters.
		parameters, err = nil, nil
	}
	if err != nil {
		return fmt.Errorf("Could not parse supplied pipeline parameters: %v.\n", err)
	}
	trigger := map[string]interface{}{"type": "manual"}
	if len(parameters) > 0 {
		trigger["parameters"] = parameters
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

	executions := make([]interface{}, 0)
	attempts := 0
	for len(executions) == 0 && attempts < 5 {
		executions, resp, err = gateClient.ExecutionsControllerApi.SearchForPipelineExecutionsByTriggerUsingGET(
			gateClient.Context,
			options.application,
			map[string]interface{}{
				"pipelineName": options.name,
				"statuses":     "RUNNING",
			})
		attempts += 1
		time.Sleep(time.Duration(attempts*attempts) * time.Second)
	}
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("Encountered an error querying pipeline execution, status code: %d\n", resp.StatusCode)
	}
	if len(executions) == 0 {
		return fmt.Errorf("Unable to start any executions, server response was: %v", resp)
	}

	refIds := make([]string, 0)
	for _, execution := range executions {
		refIds = append(refIds, execution.(map[string]interface{})["id"].(string))
	}
	util.UI.Output(fmt.Sprintf("%v", refIds))
	return nil
}
