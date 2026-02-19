package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/alexhokl/helper/authhelper"
	"github.com/alexhokl/helper/jsonhelper"
	"github.com/alexhokl/strava-cli/swagger"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

// listActivityCmd represents the list activity command
var listActivityCmd = &cobra.Command{
	Use:     "activity",
	Aliases: []string{"activities"},
	Short:   "List recent activities of the current user",
	RunE:    runListActivities,
}

type listActivityOptions struct {
	limit int32
}

var listActivityOpts listActivityOptions

func init() {
	listCmd.AddCommand(listActivityCmd)

	flags := listActivityCmd.Flags()
	flags.Int32Var(&listActivityOpts.limit, "limit", 10, "Number of activities to show")
}

func runListActivities(_ *cobra.Command, _ []string) error {
	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: savedToken.AccessToken})
	auth := context.WithValue(context.Background(), swagger.ContextOAuth2, tokenSource)
	config := swagger.NewConfiguration()
	client := swagger.NewAPIClient(config)

	activities, err := fetchActivities(client, auth, listActivityOpts.limit)
	if err != nil {
		return err
	}

	if listOpts.format == "json" {
		json, err := jsonhelper.GetJSONString(activities)
		if err != nil {
			return err
		}
		fmt.Println(json)
		return nil
	}

	var data [][]string
	for _, e := range activities {
		arr := []string{
			fmt.Sprintf("%d", e.GetId()),
			e.GetStartDate().Local().String(),
			e.GetName(),
			fmt.Sprintf("%d", e.GetTotalPhotoCount()),
		}
		data = append(data, arr)
	}

	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithRendition(tw.Rendition{Borders: tw.BorderNone}),
	)
	table.Header("ID", "Date", "Activity", "Photos")
	if err := table.Bulk(data); err != nil {
		return fmt.Errorf("failed to add table data: %w", err)
	}
	if err := table.Render(); err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}

	return nil
}

const maxPerPage int32 = 200

func fetchActivities(client *swagger.APIClient, auth context.Context, limit int32) ([]swagger.SummaryActivity, error) {
	var allActivities []swagger.SummaryActivity

	page := int32(1)
	remaining := limit

	for remaining > 0 {
		perPage := remaining
		if perPage > maxPerPage {
			perPage = maxPerPage
		}

		activities, _, err := client.ActivitiesAPI.GetLoggedInAthleteActivities(auth).
			PerPage(perPage).
			Page(page).
			Execute()
		if err != nil {
			return nil, err
		}

		allActivities = append(allActivities, activities...)

		// If we got fewer activities than requested, we've reached the end
		if int32(len(activities)) < perPage {
			break
		}

		remaining -= int32(len(activities))
		page++
	}

	return allActivities, nil
}
