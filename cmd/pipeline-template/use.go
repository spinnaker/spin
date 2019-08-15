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
	"strings"

	"github.com/spf13/cobra"
	"github.com/spinnaker/spin/util"
)

type UseOptions struct {
	*pipelineTemplateOptions
	id              string
	tag             string
	application     string
	name            string
	description     string
	variables       map[string]string
	templateType    string
	artifactAccount string
}

var (
	usePipelineTemplateShort = "Creates a pipeline configuration using a managed pipeline template"
	usePipelineTemplateLong  = "Creates a pipeline configuration using a managed pipeline template"
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
	cmd.PersistentFlags().StringVarP(&options.application, "application", "a", "", "application to get the new pipeline")
	cmd.PersistentFlags().StringVarP(&options.name, "name", "n", "", "name of the new pipeline")
	cmd.PersistentFlags().StringVarP(&options.tag, "tag", "t", "", "(optional) specific tag to query")
	cmd.PersistentFlags().StringVarP(&options.description, "description", "d", "", "(optional) description of the pipeline")
	cmd.PersistentFlags().StringVar(&options.templateType, "type", "front50/pipelineTemplate", "(optional) template type")
	cmd.PersistentFlags().StringVar(&options.artifactAccount, "artifact-acount", "front50ArtifactCredentials", "(optional) artifact account")
	cmd.PersistentFlags().StringToStringVarP(&options.variables, "variables", "v", nil, "template variables/values required by the template.  Format: key=val,key1=val1")

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

	// Build pipeline using template, output
	pipeline := buildUsingTemplate(id, options)
	util.InitUI(false, false, "")
	util.UI.JsonOutput(pipeline, util.UI.OutputFormat)

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
	templateProperty["artifactAccount"] = options.artifactAccount
	templateProperty["type"] = options.templateType
	templateProperty["reference"] = getFullTemplateID(id, options.tag)

	// Configure pipeline
	pipeline["template"] = templateProperty
	pipeline["schema"] = "v2"
	pipeline["application"] = options.application
	pipeline["name"] = options.name
	pipeline["description"] = options.description
	pipeline["variables"] = options.variables

	// Properties not supported by spin, add empty default values which can be populated manually if desired
	pipeline["exclude"] = make([]string, 0)
	pipeline["triggers"] = make([]string, 0)
	pipeline["parameters"] = make([]string, 0)
	pipeline["notifications"] = make([]string, 0)
	pipeline["stages"] = make([]string, 0)

	return pipeline
}

func getFullTemplateID(id string, tag string) string {
	// If no protocol given, add default spinnaker://
	if !strings.Contains(id, "://") {
		id = fmt.Sprintf("spinnaker://%s", id)
	}
	// Append the tag if they set one
	if tag != "" {
		id = fmt.Sprintf("%s:%s", id, tag)
	}
	// Otherwise they have set the protocol, return it back as is
	return id
}
