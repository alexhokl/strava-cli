/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/alexhokl/strava-cli/swagger"
	"github.com/antihax/optional"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listSegmentCmd = &cobra.Command{
	Use:   "segment",
	Short: "List starred segment of the current user",
	RunE:  runListSegments,
}

func init() {
	listCmd.AddCommand(listSegmentCmd)
}

func runListSegments(_ *cobra.Command, _ []string) error {
	accessToken := viper.GetString("token")
	auth := context.WithValue(context.Background(), swagger.ContextAccessToken, accessToken)
	config := swagger.NewConfiguration()
	client := swagger.NewAPIClient(config)

	opts := &swagger.SegmentsApiGetLoggedInAthleteStarredSegmentsOpts{
		PerPage: optional.NewInt32(100),
		Page:    optional.NewInt32(1),
	}
	segments, _, err := client.SegmentsApi.GetLoggedInAthleteStarredSegments(auth, opts)
	if err != nil {
		return err
	}

	var data [][]string
	for _, e := range segments {
		arr := []string{
			fmt.Sprintf("%d", e.Id),
			e.Country,
			e.Name,
			fmt.Sprintf("%.1f", e.Distance/1000.0),
			fmt.Sprintf("%.0f", e.ElevationHigh - e.ElevationLow),
			fmt.Sprintf("%.2f", (e.ElevationHigh - e.ElevationLow)/e.Distance*100),
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