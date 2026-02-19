package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/alexhokl/helper/authhelper"
	"github.com/alexhokl/helper/jsonhelper"
	"github.com/alexhokl/strava-cli/swagger"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

// listSegmentEffortCmd represents the list segment effort command
var listSegmentEffortCmd = &cobra.Command{
	Use:     "segment-effort",
	Aliases: []string{"segment-efforts"},
	Short:   "List efforts of a segment of the current user order by quickest time",
	RunE:    runListSegmentEfforts,
}

type listSegmentEffortOptions struct {
	id int32
}

var listSegmentEffortOpts listSegmentEffortOptions

func init() {
	listCmd.AddCommand(listSegmentEffortCmd)

	flags := listSegmentEffortCmd.Flags()
	flags.Int32Var(&listSegmentEffortOpts.id, "id", 0, "Segment ID")
	_ = listSegmentEffortCmd.MarkFlagRequired("id")
}

func runListSegmentEfforts(_ *cobra.Command, _ []string) error {
	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: savedToken.AccessToken})
	auth := context.WithValue(context.Background(), swagger.ContextOAuth2, tokenSource)
	config := swagger.NewConfiguration()
	client := swagger.NewAPIClient(config)

	efforts, _, err := client.SegmentEffortsAPI.GetEffortsBySegmentId(auth).
		SegmentId(listSegmentEffortOpts.id).
		StartDateLocal(time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)).
		EndDateLocal(time.Now()).
		PerPage(100).
		Execute()
	if err != nil {
		return err
	}

	if listOpts.format == "json" {
		json, err := jsonhelper.GetJSONString(efforts)
		if err != nil {
			return err
		}
		fmt.Println(json)
		return nil
	}

	var data [][]string
	for _, e := range efforts {
		duration, _ := time.ParseDuration(fmt.Sprintf("%ds", e.GetElapsedTime()))
		arr := []string{
			e.GetStartDate().Format("2006-01-02"),
			duration.String(),
			fmt.Sprintf("%.0f", e.GetAverageWatts()),
			fmt.Sprintf("%.0f", e.GetAverageCadence()),
			fmt.Sprintf("%.0f", e.GetAverageHeartrate()),
			fmt.Sprintf("%.0f", e.GetMaxHeartrate()),
		}
		data = append(data, arr)
	}

	if len(efforts) > 0 {
		segment := efforts[0].GetSegment()
		fmt.Printf("Segment: %s\n", *segment.Name)
	}

	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithRendition(tw.Rendition{Borders: tw.BorderNone}),
	)
	table.Header("Date", "Duration", "Power (W)", "Cadence", "Heart rate", "Max heart rate")
	if err := table.Bulk(data); err != nil {
		return fmt.Errorf("failed to add table data: %w", err)
	}
	if err := table.Render(); err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}

	return nil
}
