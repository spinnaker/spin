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
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/spinnaker/spin/cmd"
	"github.com/spinnaker/spin/util"
)

func TestProjectList_basic(t *testing.T) {
	ts := testGateProjectList(false)
	defer ts.Close()

	rootCmd, options := cmd.NewCmdRoot(ioutil.Discard, ioutil.Discard)
	rootCmd.AddCommand(NewProjectCmd(options))

	args := []string{"project", "list", "--gate-endpoint=" + ts.URL}
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Command failed with: %s", err)
	}
}

func TestProjectList_malformed(t *testing.T) {
	ts := testGateProjectList(true)
	defer ts.Close()

	rootCmd, options := cmd.NewCmdRoot(ioutil.Discard, ioutil.Discard)
	rootCmd.AddCommand(NewProjectCmd(options))

	args := []string{"project", "list", "--gate-endpoint=" + ts.URL}
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	if err == nil {
		t.Fatalf("Command failed with: %s", err)
	}
}

func TestProjectList_fail(t *testing.T) {
	ts := testGateFail()
	defer ts.Close()

	rootCmd, options := cmd.NewCmdRoot(ioutil.Discard, ioutil.Discard)
	rootCmd.AddCommand(NewProjectCmd(options))

	args := []string{"project", "list", "--gate-endpoint=" + ts.URL}
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	if err == nil {
		t.Fatalf("Command failed with: %s", err)
	}
}

// testGateProjectList spins up a local http server that we will configure the GateClient
// to direct requests to. When 'returnMalformed' is false, responds with a 200 and a well-formed project list.
// Returns a malformed list of project configs when 'returnMalformed' is true
func testGateProjectList(returnMalformed bool) *httptest.Server {
	mux := util.TestGateMuxWithVersionHandler()
	mux.Handle("/projects", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if returnMalformed {
			fmt.Fprintln(w, strings.TrimSpace(malformedProjectListJson))
		} else {
			fmt.Fprintln(w, strings.TrimSpace(projectListJson))
		}
	}))

	return httptest.NewServer(mux)
}

const projectListJson = `
[
 {
  "config": {
   "applications": [
    "app1"
   ],
   "clusters": [
    {
     "account": "spin-us-west-2",
     "applications": null,
     "detail": "*",
     "stack": "*"
    }
   ],
   "pipelineConfigs": [
    {
     "application": "app1",
     "pipelineConfigId": "app1-dev"
    }
   ]
  },
  "createTs": 1595273917306,
  "email": "user@spinnaker.com",
  "id": "6406b5f9-819d-495d-b1a6-3928a2a72311",
  "lastModifiedBy": "user@spinnaker.com",
  "name": "project1",
  "updateTs": 1595273918000
 }
]`

const malformedProjectListJson = `
 {
  "config": {
   "applications": [
    "app1"
   ],
   "clusters": [
    {
     "account": "spin-us-west-2",
     "applications": null,
     "detail": "*",
     "stack": "*"
    }
   ],
   "pipelineConfigs": [
    {
     "application": "app1",
     "pipelineConfigId": "app1-dev"
    }
   ]
  },
  "createTs": 1595273917306,
  "email": "user@spinnaker.com",
  "id": "6406b5f9-819d-495d-b1a6-3928a2a72311",
  "lastModifiedBy": "user@spinnaker.com",
  "name": "project1",
  "updateTs": 1595273918000
 }
]
`
