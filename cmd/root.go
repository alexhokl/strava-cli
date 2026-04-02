package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexhokl/helper/authhelper"
	"github.com/alexhokl/helper/cli"
	"github.com/alexhokl/helper/strava"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

const applicationName = "strava-cli"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               applicationName,
	Short:             "A CLI application interacting with Strava API",
	SilenceUsage:      true,
	PersistentPreRunE: validateToken,
}

func Execute() {
	_ = rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	defaultConfigDesc := fmt.Sprintf("$HOME/.config/%s/config.yaml", applicationName)
	if configDir, err := os.UserConfigDir(); err == nil {
		defaultConfigDesc = filepath.Join(configDir, applicationName, "config.yaml")
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default: %s)", defaultConfigDesc))
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	cli.ConfigureViper(cfgFile, "strava-cli", false, "")
}

func validateToken(cmd *cobra.Command, _ []string) error {
	// skips checking if it is login
	if cmd.Name() == "login" {
		return nil
	}

	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}

	config, err := getOAuthConfigurationFromViper()
	if err != nil {
		return err
	}

	if !savedToken.Valid() {
		ctx := context.Background()
		newToken, err := authhelper.RefreshToken(ctx, config.GetOAuthConfig(), savedToken)
		if err != nil {
			return fmt.Errorf("invalid token. please login again: %v", err)
		}
		_ = authhelper.SaveTokenToViper(newToken)
		return nil
	}

	return nil
}

func getOAuthConfigurationFromViper() (*authhelper.OAuthConfig, error) {
	clientID := viper.GetString("clientId")
	clientSecret := viper.GetString("clientSecret")
	if clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("client_id or client_secret is not configured")
	}

	config := &authhelper.OAuthConfig{
		ClientId:     clientID,
		ClientSecret: clientSecret,
		Scopes:       getScopes(),
		RedirectURI:  "/callback",
		Port:         port,
		Endpoint:     strava.GetOAuthEndpoint(),
	}
	return config, nil
}
