package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/alexhokl/helper/authhelper"
	"github.com/alexhokl/helper/jsonhelper"
	"github.com/alexhokl/strava-cli/swagger"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

// listSegmentCmd represents the list segment command
var listSegmentCmd = &cobra.Command{
	Use:   "segment",
	Short: "List starred segment of the current user",
	RunE:  runListSegments,
}

func init() {
	listCmd.AddCommand(listSegmentCmd)
}

func runListSegments(_ *cobra.Command, _ []string) error {
	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: savedToken.AccessToken})
	auth := context.WithValue(context.Background(), swagger.ContextOAuth2, tokenSource)
	config := swagger.NewConfiguration()
	client := swagger.NewAPIClient(config)

	segments, _, err := client.SegmentsAPI.GetLoggedInAthleteStarredSegments(auth).
		PerPage(100).
		Page(1).
		Execute()
	if err != nil {
		return err
	}

	if listOpts.format == "json" {
		json, err := jsonhelper.GetJSONString(segments)
		if err != nil {
			return err
		}
		fmt.Println(json)
		return nil
	}

	var data [][]string
	for _, e := range segments {
		arr := []string{
			fmt.Sprintf("%d", e.GetId()),
			e.GetCountry(),
			e.GetName(),
			fmt.Sprintf("%.1f", e.GetDistance()/1000.0),
			fmt.Sprintf("%.0f", e.GetElevationHigh()-e.GetElevationLow()),
			fmt.Sprintf("%.2f", (e.GetElevationHigh()-e.GetElevationLow())/e.GetDistance()*100),
		}
		data = append(data, arr)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Country", "Name", "Distance (km)", "Elevation (m)", "Average gradient (%)"})
	table.SetBorder(false)
	table.AppendBulk(data)
	table.Render()

	return nil
}
