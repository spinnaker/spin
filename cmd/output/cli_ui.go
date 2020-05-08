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

package output

import (
	"fmt"
	"io"

	"github.com/mitchellh/cli"
	"github.com/mitchellh/colorstring"
)

type UI interface {
	Success(message string)
	JSONOutput(data interface{})
	cli.Ui
}

type ColorizeUI struct {
	Colorize       *colorstring.Colorize
	OutputColor    string
	InfoColor      string
	ErrorColor     string
	WarnColor      string
	SuccessColor   string
	UI             cli.Ui
	Quiet          bool
	OutputFormater OutputFormater
}

func NewUI(
	quiet, color bool,
	outputFormater OutputFormater,
	outWriter, errWriter io.Writer,
) *ColorizeUI {
	return &ColorizeUI{
		Colorize: &colorstring.Colorize{
			Colors:  colorstring.DefaultColors,
			Disable: !color,
			Reset:   true,
		},
		ErrorColor:   "[red]",
		WarnColor:    "[yellow]",
		InfoColor:    "[blue]",
		SuccessColor: "[bold][green]",
		UI: &cli.BasicUi{
			Writer:      outWriter,
			ErrorWriter: errWriter,
		},
		Quiet:          quiet,
		OutputFormater: outputFormater,
	}
}

func (u *ColorizeUI) Ask(query string) (string, error) {
	return u.UI.Ask(u.colorize(query, u.OutputColor))
}

func (u *ColorizeUI) AskSecret(query string) (string, error) {
	return u.UI.AskSecret(u.colorize(query, u.OutputColor))
}

func (u *ColorizeUI) Output(message string) {
	u.UI.Output(u.colorize(message, u.OutputColor))
}

// JSONOutput prints the data specified using the configured OutputFormater.
func (u *ColorizeUI) JSONOutput(data interface{}) {
	output, err := u.OutputFormater(data)
	if err != nil {
		u.Error(fmt.Sprintf("%v", err))
	}
	u.Output(string(output))
}

func (u *ColorizeUI) Success(message string) {
	if !u.Quiet {
		u.UI.Info(u.colorize(message, u.SuccessColor))
	}
}

func (u *ColorizeUI) Info(message string) {
	if !u.Quiet {
		u.UI.Info(u.colorize(message, u.InfoColor))
	}
}

func (u *ColorizeUI) Error(message string) {
	u.UI.Error(u.colorize(message, u.ErrorColor))
}

func (u *ColorizeUI) Warn(message string) {
	if !u.Quiet {
		u.UI.Warn(u.colorize(message, u.WarnColor))
	}
}

func (u *ColorizeUI) colorize(message string, color string) string {
	if color == "" {
		return message
	}

	return u.Colorize.Color(fmt.Sprintf("%s%s", color, message))
}
