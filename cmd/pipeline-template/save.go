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

package pipeline_template

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spinnaker/spin/cmd/gateclient"
	"github.com/spinnaker/spin/util"
	"net/http"
)

type SaveOptions struct {
	*pipelineTemplateOptions
	output       string
	templateFile string
	tag string
}

var (
	saveTemplateShort   = "Save the provided pipeline template"
	saveTemplateLong    = "Save the provided pipeline template"
)

func NewSaveCmd(pipelineTemplateOptions pipelineTemplateOptions) *cobra.Command {
	options := SaveOptions{
		pipelineTemplateOptions: &pipelineTemplateOptions,
	}
	cmd := &cobra.Command{
		Use:     "save",
		Aliases: []string{},
		Short:   saveTemplateShort,
		Long:    saveTemplateLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			return savePipelineTemplate(cmd, options)
		},
	}

	cmd.PersistentFlags().StringVarP(&options.templateFile, "file",
		"f", "", "path to the pipeline template file")
	cmd.PersistentFlags().StringVar(&options.tag, "tag", "",
		"(optional) specific tag to tag pipeline template with")

	return cmd
}

func savePipelineTemplate(cmd *cobra.Command, options SaveOptions) error {
	gateClient, err := gateclient.NewGateClient(cmd.InheritedFlags())
	if err != nil {
		return err
	}

	templateJson, err := util.ParseJsonFromFileOrStdin(options.templateFile, false)
	if err != nil {
		return err
	}

	valid := true
	if _, exists := templateJson["id"]; !exists {
		util.UI.Error("Required pipeline template key 'id' missing...\n")
		valid = false
	}
	if _, exists := templateJson["schema"]; !exists {
		util.UI.Error("Required pipeline template key 'schema' missing...\n")
		valid = false
	}
	if !valid {
		return fmt.Errorf("Submitted pipeline is invalid: %s\n", templateJson)
	}

	templateId := templateJson["id"].(string)

	queryParams := map[string]interface{}{}
	if options.tag != "" {
		queryParams["tag"] = options.tag
	}

	_, resp, queryErr := gateClient.V2PipelineTemplatesControllerApi.GetUsingGET2(gateClient.Context, templateId, queryParams)

	var saveResp *http.Response
	var saveErr error
	if resp.StatusCode == http.StatusOK {
		saveResp, saveErr = gateClient.V2PipelineTemplatesControllerApi.UpdateUsingPOST1(gateClient.Context, templateId, templateJson, queryParams)
	} else if resp.StatusCode == http.StatusNotFound {
		saveResp, saveErr = gateClient.V2PipelineTemplatesControllerApi.CreateUsingPOST1(gateClient.Context, templateJson, queryParams)
	} else {
		if queryErr != nil {
      return queryErr
		}
		return fmt.Errorf("Encountered an unexpected status code %d querying pipeline template with id %s\n",
			resp.StatusCode, templateId)
	}

	if saveErr != nil {
    return saveErr
	}

	if saveResp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Encountered an error saving pipeline template %v, status code: %d\n",
			templateJson,
			saveResp.StatusCode)
	}

	util.UI.Info(util.Colorize().Color(fmt.Sprintf("[reset][bold][green]Pipeline template save succeeded")))
	return nil
}
