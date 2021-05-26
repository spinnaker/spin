package config

import (
	"testing"

	"sigs.k8s.io/yaml"
)

func TestBasicMarshalling(t *testing.T) {
	var cfg Config

	cfg.Gate.Endpoint = "test"

	bs, err := yaml.Marshal(cfg)

	if err != nil {
		t.Errorf("Unexpected error unmarshalling YAML: %v", err)
	}

	got := string(bs)
	want := "auth: null\ngate:\n  endpoint: test\n"

	if got != want {
		t.Errorf("Unexpected marshalled YAML, saw '%s', want '%s'", got, want)
	}
}
