package config

import (
	"github.com/spinnaker/spin/config/authconfig"
)

type Config struct {
	Auth authconfig.AuthConfig `yaml:"auth"`
}
