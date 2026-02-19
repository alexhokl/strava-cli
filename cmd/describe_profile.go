package cmd

import (
	"context"
	"fmt"

	"github.com/alexhokl/helper/authhelper"
	"github.com/alexhokl/strava-cli/swagger"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

// describeProfileCmd represents the profile command
var describeProfileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Describe profile of the current user",
	RunE:  runDescribeProfile,
}

func init() {
	describeCmd.AddCommand(describeProfileCmd)
}

func runDescribeProfile(_ *cobra.Command, _ []string) error {
	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: savedToken.AccessToken})
	auth := context.WithValue(context.Background(), swagger.ContextOAuth2, tokenSource)
	config := swagger.NewConfiguration()
	client := swagger.NewAPIClient(config)
	profile, _, err := client.AthletesAPI.GetLoggedInAthlete(auth).Execute()
	if err != nil {
		return err
	}
	fmt.Printf(
		"Hi %s!\nYour FTP is %dW with weight %fkg\n",
		profile.GetFirstname(),
		profile.GetFtp(),
		profile.GetWeight(),
	)
	return nil
}
