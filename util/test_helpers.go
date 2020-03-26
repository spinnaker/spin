package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

func TestGateMuxWithVersionHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/version", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := map[string]string{
			"version": "Unknown",
		}
		b, _ := json.Marshal(&payload)
		fmt.Fprintln(w, string(b))
	}))

	return mux
}

func NewRootCmdForTest() *cobra.Command {
	rootCmd := &cobra.Command{}
	rootCmd.PersistentFlags().String("config", "", "config file (default is $HOME/.spin/config)")
	rootCmd.PersistentFlags().String("gate-endpoint", "", "Gate (API server) endpoint. Default http://localhost:8084")
	rootCmd.PersistentFlags().Bool("insecure", false, "Ignore Certificate Errors")
	rootCmd.PersistentFlags().Bool("quiet", false, "Squelch non-essential output")
	rootCmd.PersistentFlags().Bool("no-color", false, "Disable color")
	rootCmd.PersistentFlags().String("output", "", "Configure output formatting")
	rootCmd.PersistentFlags().String("default-headers", "", "Configure additional headers for gate client requests")
	InitUI(false, false, "")
	return rootCmd
}

// ExecCmdForTest executes the command and returns the STDOUT as a string.
func ExecCmdForTest(cmd *cobra.Command) (string, error) {
	buffer := bytes.NewBufferString("")
	cmd.SetOut(buffer)
	err := cmd.Execute()
	out, bufferErr := ioutil.ReadAll(buffer)
	if bufferErr != nil {
		return "", fmt.Errorf("Failed to read command output buffer: %v", err)
	}
	return string(out), err
}
