package main

import (
	"github.com/mitchellh/cli"
	"github.com/spinnaker/spin/command"
	"github.com/spinnaker/spin/command/pipelines"
)

var Commands map[string]cli.CommandFactory

func init() {
	meta := command.ApiMeta{}

	Commands = map[string]cli.CommandFactory{
		"pipeline save": func() (cli.Command, error) {
			return &pipelines.PipelineSaveCommand{
				ApiMeta: meta,
			}, nil
		},
		"pipeline list": func() (cli.Command, error) {
			return &pipelines.PipelineListCommand{
				ApiMeta: meta,
			}, nil
		},
	}
}
