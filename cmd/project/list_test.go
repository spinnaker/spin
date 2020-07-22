// Copyright (c) 2020, Anosua Chini Mukhopadhyay
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

package project

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/spinnaker/spin/cmd"
	"github.com/spinnaker/spin/util"
)

// const (
// 	NAME              = "prj"
// 	EMAIL             = "user@spinnaker.com"
// 	PROJECT_TEST_FILE = "../../util/json_test_files/example_project.json"
// )

func TestProjectList_basic(t *testing.T) {
	saveBuffer := new(bytes.Buffer)
	ts := testGateProjectListSuccess(saveBuffer)
	defer ts.Close()

	rootCmd, options := cmd.NewCmdRoot(ioutil.Discard, ioutil.Discard)
	rootCmd.AddCommand(NewProjectCmd(options))

	args := []string{
		"project", "list",
		"--gate-endpoint=" + ts.URL,
	}
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Command failed with: %s", err)
	}

	expected := strings.TrimSpace(testProjectTaskJsonStr)
	recieved := saveBuffer.Bytes()
	util.TestPrettyJsonDiff(t, "list request body", expected, recieved)
}

// func TestProjectList_fail(t *testing.T) {
// 	ts := testGateFail()
// 	defer ts.Close()

// 	rootCmd, options := cmd.NewCmdRoot(ioutil.Discard, ioutil.Discard)
// 	rootCmd.AddCommand(NewProjectCmd(options))

// 	args := []string{
// 		"project", "list",
// 		"--gate-endpoint=" + ts.URL,
// 	}
// 	rootCmd.SetArgs(args)
// 	err := rootCmd.Execute()
// 	if err == nil {
// 		t.Fatalf("Command failed with: %s", err)
// 	}
// }

// func TestProjectList_flags(t *testing.T) {
// 	saveBuffer := new(bytes.Buffer)
// 	ts := testGateProjectListSuccess(saveBuffer)
// 	defer ts.Close()

// 	rootCmd, options := cmd.NewCmdRoot(ioutil.Discard, ioutil.Discard)
// 	rootCmd.AddCommand(NewProjectCmd(options))

// 	args := []string{
// 		"project", "list",
// 		"--gate-endpoint=" + ts.URL,
// 	}
// 	rootCmd.SetArgs(args)
// 	err := rootCmd.Execute()
// 	if err == nil {
// 		t.Fatalf("Command failed with: %s", err)
// 	}

// 	expected := ""
// 	recieved := strings.TrimSpace(saveBuffer.String())
// 	if expected != recieved {
// 		t.Fatalf("Unexpected list request body:\n%s", recieved)
// 	}
// }

// func TestProjectList_filejson(t *testing.T) {
// 	saveBuffer := new(bytes.Buffer)
// 	ts := testGateProjectListSuccess(saveBuffer)
// 	defer ts.Close()

// 	rootCmd, options := cmd.NewCmdRoot(ioutil.Discard, ioutil.Discard)
// 	rootCmd.AddCommand(NewProjectCmd(options))

// 	args := []string{
// 		"project", "list",
// 		"--gate-endpoint", ts.URL,
// 	}
// 	rootCmd.SetArgs(args)
// 	err := rootCmd.Execute()
// 	if err != nil {
// 		t.Fatalf("Command failed with: %s", err)
// 	}

// 	expected := strings.TrimSpace(testExpandedProjectTaskJsonStr)
// 	recieved := saveBuffer.Bytes()
// 	util.TestPrettyJsonDiff(t, "list request body", expected, recieved)
// }

// testGateProjectListSuccess spins up a local http server that we will configure the GateClient
// to direct requests to. Responds with successful responses to pipeline execute API calls.
// Writes request body to buffer for testing.
func testGateProjectListSuccess(buffer io.Writer) *httptest.Server {
	mux := util.TestGateMuxWithVersionHandler()
	mux.Handle(
		"/tasks",
		util.NewTestBufferHandlerFunc(http.MethodPost, buffer, http.StatusOK, strings.TrimSpace(testAppTaskRefJsonStr)),
	)
	mux.Handle("/tasks/id", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, strings.TrimSpace(testProjectTaskStatusJsonStr))
	}))
	return httptest.NewServer(mux)
}

// const testAppTaskRefJsonStr = `
// {
//  "ref": "/tasks/id"
// }
// `
// const testProjectTaskStatusJsonStr = `
// {
//  "status": "SUCCEEDED"
// }
// `

// const testExpandedProjectTaskJsonStr = `
// {
//  "application": "spinnaker",
//  "description": "Create Project: example project",
//  "job": [
//   {
//    "project": {
//     "config": {
//      "applications": [],
//      "clusters": [],
//      "pipelineConfigs": []
//     },
//     "email": "user@spinnaker.com",
//     "name": "example project",
//     "user": "user@spinnaker.com"
//    },
//    "type": "upsertProject",
//    "user": "user@spinnaker.com"
//   }
//  ],
//  "project": "example project"
// }
// `

// const testProjectTaskJsonStr = `
// {
//  "application": "spinnaker",
//  "description": "Create Project: prj",
//  "job": [
//   {
//    "project": {
//     "email": "user@spinnaker.com",
//     "name": "prj"
//    },
//    "type": "upsertProject",
//    "user": "user@spinnaker.com"
//   }
//  ],
//  "project": "prj"
// }
// `
