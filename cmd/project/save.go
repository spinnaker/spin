package project

import (
	"errors"
	"fmt"
    "net/http"

	"github.com/spf13/cobra"
	orca_tasks "github.com/spinnaker/spin/cmd/orca-tasks"
	"github.com/spinnaker/spin/util"
)

type saveOptions struct {
	*projectOptions
	projectFile string
	projectName string
	ownerEmail  string
}

var (
	saveProjectShort = "Save the provided project"
	saveOrojectLong  = "Save the specified project"
)

func NewSaveCmd(prjOptions *projectOptions) *cobra.Command {
	options := &saveOptions{
		projectOptions: prjOptions,
	}
	cmd := &cobra.Command{
		Use:   "save",
		Short: saveProjectShort,
		Long:  saveOrojectLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			return saveProject(cmd, options)
		},
	}
	cmd.PersistentFlags().StringVarP(&options.projectFile, "file", "f", "", "path to the project file")
	cmd.PersistentFlags().StringVarP(&options.projectName, "name", "n", "", "name of the project")
	cmd.PersistentFlags().StringVarP(&options.ownerEmail, "email", "e", "", "email of the project owner")

	return cmd
}

func saveProject(cmd *cobra.Command, options *saveOptions) error {
	// TODO: check for existing project

	initialProject, err := util.ParseJsonFromFileOrStdin(options.projectFile, true)
	if err != nil {
		return fmt.Errorf("Could not parse supplied project: %v.\n", err)
	}

	var project map[string]interface{}
	if initialProject != nil && len(initialProject) > 0 {
		project = initialProject
		if options.projectName != "" {
			options.Ui.Warn("Overriding project name with explicit flag values.\n")
			project["name"] = options.projectName
		}
		if options.ownerEmail != "" {
			options.Ui.Warn("Overriding project owner email with explicit flag values.\n")
			project["email"] = options.ownerEmail
		}
	} else {
		if options.projectName == "" || options.ownerEmail == "" {
			return errors.New("Required project parameters missing, exiting...")
		}
		project = map[string]interface{}{
			"name":  options.projectName,
			"email": options.ownerEmail,
		}
	}

    projectName := fmt.Sprintf("%s", project["name"])
	id, err := doesProjectExist(projectName, options)
    if id != "" {
        project["id"] = id
    }

	createProjectTask := map[string]interface{}{
		"job":         []interface{}{map[string]interface{}{"type": "upsertProject", "project": project, "user": project["email"]}},
		"application": "spinnaker",
		"project":     projectName,
		"description": fmt.Sprintf("Create Project: %s", projectName),
	}

	ref, _, err := options.GateClient.TaskControllerApi.TaskUsingPOST1(options.GateClient.Context, createProjectTask)
	if err != nil {
		return err
	}

	err = orca_tasks.WaitForSuccessfulTask(options.GateClient, ref, 5)
	if err != nil {
		return err
	}

	options.Ui.Success("Project save succeeded")
	return nil
}

func doesProjectExist(projectName string, options *saveOptions) (string, error) {
	project, resp, err := options.GateClient.ProjectControllerApi.GetUsingGET1(options.GateClient.Context, projectName)
	if resp != nil {
		if resp.StatusCode == http.StatusNotFound {
			return "", nil
		} else if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("Encountered an error getting project, status code: %d\n", resp.StatusCode)
		} else {
			return fmt.Sprintf("%s", project["id"]), nil
		}
	}

	return "", err
}
