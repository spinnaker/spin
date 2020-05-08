// Copyright (c) 2019, Waze, Inc.
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

package canary_config

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spinnaker/spin/util"
)

type saveOptions struct {
	*canaryConfigOptions
	templateFile string
}

const (
	saveTemplateShort = "Save the provided canary config"
	saveTemplateLong  = "Save the provided canary config"
)

func NewSaveCmd(canaryConfigOptions *canaryConfigOptions) *cobra.Command {
	options := &saveOptions{
		canaryConfigOptions: canaryConfigOptions,
	}
	cmd := &cobra.Command{
		Use:     "save",
		Aliases: []string{},
		Short:   saveTemplateShort,
		Long:    saveTemplateLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			return saveCanaryConfig(options)
		},
	}

	cmd.PersistentFlags().StringVarP(&options.templateFile, "file",
		"f", "", "path to the canary config file")

	return cmd
}

func saveCanaryConfig(options *saveOptions) error {
	templateJSON, err := util.ParseJSONFromFileOrStdin(options.templateFile, false)
	if err != nil {
		return err
	}

	if _, exists := templateJSON["id"]; !exists {
		options.UI.Error("Required canary config key 'id' missing...\n")
		return fmt.Errorf("Submitted canary config is invalid: %s\n", templateJSON)
	}

	templateID := templateJSON["id"].(string)

	_, resp, queryErr := options.GateClient.V2CanaryConfigControllerApi.GetCanaryConfigUsingGET(
		options.GateClient.Context, templateID, map[string]interface{}{})
	if queryErr != nil {
		return queryErr
	}
	var saveResp *http.Response
	var saveErr error
	switch {
	case resp.StatusCode == http.StatusOK:
		_, saveResp, saveErr = options.GateClient.V2CanaryConfigControllerApi.UpdateCanaryConfigUsingPUT(
			options.GateClient.Context, templateJSON, templateID, map[string]interface{}{})
	case resp.StatusCode == http.StatusNotFound:
		_, saveResp, saveErr = options.GateClient.V2CanaryConfigControllerApi.CreateCanaryConfigUsingPOST(
			options.GateClient.Context, templateJSON, map[string]interface{}{})
	default:
		return fmt.Errorf(
			"Encountered an unexpected status code %d querying canary config with id %s\n",
			resp.StatusCode, templateID)
	}
	if saveErr != nil {
		return saveErr
	}

	if saveResp.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"Encountered an error saving canary config %v, status code: %d\n",
			templateJSON, saveResp.StatusCode)
	}

	options.UI.Success("Canary config save succeeded")
	return nil
}
