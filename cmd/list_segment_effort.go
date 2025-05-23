package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/alexhokl/helper/authhelper"
	"github.com/alexhokl/helper/jsonhelper"
	"github.com/alexhokl/strava-cli/swagger"
	"github.com/antihax/optional"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listSegmentEffortCmd represents the list segment effort command
var listSegmentEffortCmd = &cobra.Command{
	Use:   "segment-effort",
	Short: "List efforts of a segment of the current user order by quickest time",
	RunE:  runListSegmentEfforts,
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
	auth := context.WithValue(context.Background(), swagger.ContextAccessToken, savedToken.AccessToken)
	config := swagger.NewConfiguration()
	client := swagger.NewAPIClient(config)

	opts := &swagger.SegmentEffortsApiGetEffortsBySegmentIdOpts{
		StartDateLocal: optional.NewTime(time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)),
		EndDateLocal:   optional.NewTime(time.Now()),
		PerPage:        optional.NewInt32(100),
	}
	efforts, _, err := client.SegmentEffortsApi.GetEffortsBySegmentId(auth, listSegmentEffortOpts.id, opts)
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
		duration, _ := time.ParseDuration(fmt.Sprintf("%ds", e.ElapsedTime))
		arr := []string{
			e.StartDate.Format("2006-01-02"),
			duration.String(),
			fmt.Sprintf("%.0f", e.AverageWatts),
			fmt.Sprintf("%.0f", e.AverageCadence),
			fmt.Sprintf("%.0f", e.AverageHeartrate),
			fmt.Sprintf("%.0f", e.MaxHeartrate),
		}
		data = append(data, arr)
	}

	if len(efforts) > 0 {
		fmt.Printf("Segment: %s\n", efforts[0].Segment.Name)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Duration", "Power (W)", "Cadence", "Heart rate", "Max heart rate"})
	table.SetBorder(false)
	table.AppendBulk(data)
	table.Render()

	return nil
}
