// Copyright 2019 New Relic Corporation. All rights reserved.
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

package account

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type listOptions struct {
	*accountOptions
	expand bool
}

var (
	listAccountShort   = "List the all accounts"
	listAccountLong    = "List the all accounts"
	listAccountExample = "usage: spin account list [options]"
)

func NewListCmd(accOptions *accountOptions) *cobra.Command {
	options := &listOptions{
		accountOptions: accOptions,
		expand:         false,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   listAccountShort,
		Long:    listAccountLong,
		Example: listAccountExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listAccount(options)
		},
	}
	return cmd
}

func listAccount(options *listOptions) error {
	accountList, resp, err := options.GateClient.CredentialsControllerApi.GetAccountsUsingGET(options.GateClient.Context, map[string]interface{}{"expand": options.expand})
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Encountered an error listing accounts, status code: %d\n", resp.StatusCode)
	}

	options.Ui.JsonOutput(accountList)
	return nil
}
