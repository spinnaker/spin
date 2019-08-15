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
	"os"
	"testing"
)

var testAppName = "test-application"
var testPipelineName = "test-pipeline"
var testDescription = "test-description"
var testVariables = "one=1,two=2,three=3"

func TestPipelineTemplateUse_basic(t *testing.T) {
	args := []string{"pipeline-template", "use", "test-template-id", "--application", testAppName,
		"--name", testPipelineName,
		"--description", testDescription,
		fmt.Sprintf("--variables=%s", testVariables)}

	currentCmd := NewUseCmd(pipelineTemplateOptions{})
	rootCmd := getRootCmdForTest()
	pipelineTemplateCmd := NewPipelineTemplateCmd(os.Stdout)
	pipelineTemplateCmd.AddCommand(currentCmd)
	rootCmd.AddCommand(pipelineTemplateCmd)

	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Command failed with: %s", err)
	}
}

func TestPipelineTemplateUse_basicShort(t *testing.T) {
	args := []string{"pipeline-template", "use", "test-template-id", "-a", testAppName,
		"-n", testPipelineName,
		"-d", testDescription,
		fmt.Sprintf("-v=%s", testVariables)}

	currentCmd := NewUseCmd(pipelineTemplateOptions{})
	rootCmd := getRootCmdForTest()
	pipelineTemplateCmd := NewPipelineTemplateCmd(os.Stdout)
	pipelineTemplateCmd.AddCommand(currentCmd)
	rootCmd.AddCommand(pipelineTemplateCmd)

	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Command failed with: %s", err)
	}
}

func TestPipelineTemplateUse_missingFlags(t *testing.T) {
	args := []string{"pipeline-template", "use"} // Missing id, application, name
	currentCmd := NewUseCmd(pipelineTemplateOptions{})
	rootCmd := getRootCmdForTest()
	pipelineTemplateCmd := NewPipelineTemplateCmd(os.Stdout)
	pipelineTemplateCmd.AddCommand(currentCmd)
	rootCmd.AddCommand(pipelineTemplateCmd)

	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	if err == nil {
		t.Fatalf("Expected failure but command succeeded")
	}
}
