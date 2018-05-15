package command

import (
	"flag"
	"os"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/mitchellh/colorstring"
	gate "github.com/spinnaker/spin/gateapi"
)

// ApiMeta is the state & utility shared by our commands.
type ApiMeta struct {
	// The exported fields below should be set by anyone using a command
	// with an ApiMeta field. These are expected to be set externally
	// (not from within the command itself).

	Color bool   // True if output should be colored
	Ui    cli.Ui // Ui for output

	// Internal fields
	color bool

	// This is the set of flags global to the command parser.
	gateEndpoint string

	// Gate Api client.
	GateClient *gate.APIClient
}

// GlobalFlagSet adds all global options to the flagset, and returns the flagset object
// for further modification by the subcommand.
func (m *ApiMeta) GlobalFlagSet(cmd string) *flag.FlagSet {
	f := flag.NewFlagSet(cmd, flag.ContinueOnError)

	f.StringVar(&m.gateEndpoint, "gate-endpoint", "http://localhost:8084",
		"Gate (API server) endpoint")

	f.Usage = func() {}

	return f
}

// process with process the meta-parameters out of the arguments. This
// potentially modifies the args in-place. It will return the resulting slice.
func (m *ApiMeta) Process(args []string) ([]string, error) {
	// Api client initialization.
	cfg := &gate.Configuration{
		BasePath:      "http://localhost:8084",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Spin CLI version",
	}
	m.GateClient = gate.NewAPIClient(cfg)

	// Colorization.
	m.Color = true
	m.color = m.Color
	for i, v := range args {
		if v == "--no-color" {
			m.color = false
			m.Color = false
			args = append(args[:i], args[i+1:]...)
			break
		}
	}

	// Set the Ui
	m.Ui = &ColorizeUi{
		Colorize:   m.Colorize(),
		ErrorColor: "[red]",
		WarnColor:  "[yellow]",
		InfoColor:  "[blue]",
		Ui:         &cli.BasicUi{Writer: os.Stdout},
	}
	return args, nil
}

// Colorize initializes the ui colorization.
func (m *ApiMeta) Colorize() *colorstring.Colorize {
	return &colorstring.Colorize{
		Colors:  colorstring.DefaultColors,
		Disable: !m.color,
		Reset:   true,
	}
}

func (m *ApiMeta) Help() string {
	help := `
Global Options:

	--gate-endpoint         Gate (API server) endpoint.
	`

	return strings.TrimSpace(help)
}
