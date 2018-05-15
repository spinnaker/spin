package authconfig

import (
	"github.com/spinnaker/spin/config/authconfig/x509"
)

type AuthConfig struct {
	Enabled bool            `yaml:"enabled"`
	X509    x509.X509Config `yaml:"x509"`
}
