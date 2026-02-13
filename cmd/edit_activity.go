package cmd

import (
	"context"
	"fmt"

	"github.com/alexhokl/helper/authhelper"
	"github.com/alexhokl/strava-cli/swagger"
	"github.com/alexhokl/strava-cli/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

type editActivityOptions struct {
	id int64
}

var editActivityOpts editActivityOptions

// editActivityCmd represents the edit activity command
var editActivityCmd = &cobra.Command{
	Use:   "activity",
	Short: "Edit an activity",
	RunE:  runEditActivities,
}

func init() {
	editCmd.AddCommand(editActivityCmd)

	flags := editActivityCmd.Flags()
	flags.Int64Var(&editActivityOpts.id, "id", 0, "Activity ID")
	_ = editActivityCmd.MarkFlagRequired("id")
}

func runEditActivities(_ *cobra.Command, _ []string) error {
	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: savedToken.AccessToken})
	auth := context.WithValue(context.Background(), swagger.ContextOAuth2, tokenSource)
	config := swagger.NewConfiguration()
	client := swagger.NewAPIClient(config)

	activity, _, err := client.ActivitiesAPI.GetActivityById(auth, editActivityOpts.id).
		IncludeAllEfforts(false).
		Execute()
	if err != nil {
		return err
	}

	p := tea.NewProgram(ui.NewEditorModel(activity.GetName(), activity.GetDescription()))
	updatedModel, err := p.Run()
	if err != nil {
		return err
	}
	updatedEditorModel := updatedModel.(ui.EditorModel)
	if !updatedEditorModel.HasUpdate() {
		fmt.Println("No update made")
		return nil
	}

	updatableActivity := swagger.UpdatableActivity{}
	updatableActivity.SetCommute(activity.GetCommute())
	updatableActivity.SetTrainer(activity.GetTrainer())
	updatableActivity.SetHideFromHome(activity.GetHideFromHome())
	updatableActivity.SetDescription(updatedEditorModel.Description())
	updatableActivity.SetName(updatedEditorModel.Name())
	updatableActivity.SetType(activity.GetType())
	updatableActivity.SetSportType(activity.GetSportType())
	updatableActivity.SetGearId(activity.GetGearId())

	_, _, err = client.ActivitiesAPI.UpdateActivityById(auth, editActivityOpts.id).
		Body(updatableActivity).
		Execute()
	if err != nil {
		return err
	}

	fmt.Println("Updated completed.")

	return nil
}
