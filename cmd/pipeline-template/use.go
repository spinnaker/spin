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
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spinnaker/spin/cmd/gateclient"
	"github.com/spinnaker/spin/util"
)

type UseOptions struct {
	*pipelineTemplateOptions
	id                  string
	tag                 string
	application         string
	name                string
	description         string
	disableServerChecks bool
}

var (
	usePipelineTemplateShort = "Generates a pipeline JSON configuration using the specified pipeline template"
	usePipelineTemplateLong  = "Generates a pipeline JSON configuration using the specified pipeline template"
)

func NewUseCmd(pipelineTemplateOptions pipelineTemplateOptions) *cobra.Command {
	options := UseOptions{
		pipelineTemplateOptions: &pipelineTemplateOptions,
	}
	cmd := &cobra.Command{
		Use:   "use",
		Short: usePipelineTemplateShort,
		Long:  usePipelineTemplateLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			return usePipelineTemplate(cmd, options, args)
		},
	}

	cmd.PersistentFlags().StringVar(&options.id, "id", "", "id of the pipeline template")
	cmd.PersistentFlags().StringVar(&options.application, "application", "", "application to get the new pipeline")
	cmd.PersistentFlags().StringVar(&options.name, "name", "", "name of the new pipeline")
	cmd.PersistentFlags().StringVar(&options.tag, "tag", "", "(optional) specific tag to query")
	cmd.PersistentFlags().StringVar(&options.description, "description", "", "(optional) description of the pipeline")
	cmd.PersistentFlags().BoolVar(&options.disableServerChecks, "disable-server-checks", false, "(optional) disables checking if the app and template exist")

	return cmd
}

func usePipelineTemplate(cmd *cobra.Command, options UseOptions, args []string) error {
	id, errID := getTemplateID(options, args)
	if errID != nil {
		return errID
	}

	// Check required params
	options.application = strings.TrimSpace(options.application)
	if options.application == "" {
		return errors.New("no application name supplied, exiting")
	}

	options.name = strings.TrimSpace(options.name)
	if options.name == "" {
		return errors.New("no pipeline name supplied, exiting")
	}

	// Check if PT and app exist
	gateClient, err := gateclient.NewGateClient(cmd.InheritedFlags())
	if !options.disableServerChecks {
		if err != nil {
			return err
		}

		errExists := checkExists(gateClient, id, options)
		if errExists != nil {
			return errExists
		}
	}

	pipeline := buildUsingTemplate(id, options)
	util.UI.JsonOutput(pipeline, nil)
	return nil
}

func getTemplateID(options UseOptions, args []string) (string, error) {
	// Check options if they passed in like --id
	optionsID := strings.TrimSpace(options.id)
	if optionsID != "" {
		return optionsID, nil
	}
	// Otherwise get from arguments
	argsID, err := util.ReadArgsOrStdin(args)
	if err != nil {
		return "", err
	}
	argsID = strings.TrimSpace(argsID)
	if argsID == "" {
		return "", errors.New("no pipeline template id supplied, exiting")
	}

	return argsID, nil
}

func buildUsingTemplate(id string, options UseOptions) map[string]interface{} {
	pipeline := make(map[string]interface{})
	templateProperty := make(map[string]interface{})

	// Configure pipeline.template
	templateProperty["artifactAccount"] = "front50ArtifactCredentials"
	templateProperty["type"] = fmt.Sprintf("spinnaker://%s", id)
	templateProperty["reference"] = "front50ArtifactCredentials"

	// Configure pipeline
	pipeline["template"] = templateProperty
	pipeline["schema"] = "v2"
	pipeline["application"] = options.application
	pipeline["name"] = options.name
	pipeline["description"] = options.description

	// Properties not supported by spin, add empty default values which can be populated manually if desired
	pipeline["exclude"] = make([]string, 0)
	pipeline["triggers"] = make([]string, 0)
	pipeline["parameters"] = make([]string, 0)
	pipeline["notifications"] = make([]string, 0)
	pipeline["stages"] = make([]string, 0)

	return pipeline
}

func checkExists(gateClient *gateclient.GatewayClient, id string, options UseOptions) error {
	_, respApp, errApp := gateClient.ApplicationControllerApi.GetApplicationUsingGET(gateClient.Context,
		options.application, nil)
	if errApp != nil {
		return errApp
	}

	if respApp.StatusCode != http.StatusOK {
		return fmt.Errorf("encountered an error getting application with name %s, status code: %d",
			options.application,
			respApp.StatusCode)
	}

	// Check template exists
	queryParams := map[string]interface{}{}
	if options.tag != "" {
		queryParams["tag"] = options.tag
	}
	_, respPT, errPT := gateClient.V2PipelineTemplatesControllerApi.GetUsingGET2(gateClient.Context,
		id, queryParams)

	if errPT != nil {
		return errPT
	}

	if respPT.StatusCode != http.StatusOK {
		return fmt.Errorf("encountered an error getting pipeline template with id %s, status code: %d",
			id,
			respPT.StatusCode)
	}

	return nil
}
