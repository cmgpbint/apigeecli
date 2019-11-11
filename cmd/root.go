package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/srinandan/apigeecli/cmd/apis"
	"github.com/srinandan/apigeecli/cmd/apps"
	cache "github.com/srinandan/apigeecli/cmd/cache"
	"github.com/srinandan/apigeecli/cmd/developers"
	"github.com/srinandan/apigeecli/cmd/env"
	flowhooks "github.com/srinandan/apigeecli/cmd/flowhooks"
	"github.com/srinandan/apigeecli/cmd/iam"
	"github.com/srinandan/apigeecli/cmd/keyaliases"
	"github.com/srinandan/apigeecli/cmd/keystores"
	"github.com/srinandan/apigeecli/cmd/kvm"
	"github.com/srinandan/apigeecli/cmd/org"
	"github.com/srinandan/apigeecli/cmd/products"
	"github.com/srinandan/apigeecli/cmd/projects"
	res "github.com/srinandan/apigeecli/cmd/res"
	"github.com/srinandan/apigeecli/cmd/shared"
	"github.com/srinandan/apigeecli/cmd/sharedflows"
	"github.com/srinandan/apigeecli/cmd/sync"
	targetservers "github.com/srinandan/apigeecli/cmd/targetservers"
	"github.com/srinandan/apigeecli/cmd/token"
)

//RootCmd to manage apigeecli
var RootCmd = &cobra.Command{
	Use:   "apigeecli",
	Short: "Utility to work with Apigee APIs.",
	Long:  "This command lets you interact with Apigee APIs.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		shared.Init()
		return shared.SetAccessToken()
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().BoolVarP(&shared.RootArgs.LogInfo, "log", "l",
		false, "Log Information")

	RootCmd.PersistentFlags().StringVarP(&shared.RootArgs.Token, "token", "t",
		"", "Google OAuth Token")
	_ = viper.BindPFlag("token", RootCmd.PersistentFlags().Lookup("token"))

	RootCmd.PersistentFlags().StringVarP(&shared.RootArgs.ServiceAccount, "account", "a",
		"", "Path Service Account private key in JSON")
	_ = viper.BindPFlag("account", RootCmd.PersistentFlags().Lookup("account"))

	RootCmd.PersistentFlags().BoolVar(&shared.RootArgs.SkipCache, "skipCache",
		false, "Skip caching Google OAuth Token")

	RootCmd.PersistentFlags().BoolVar(&shared.RootArgs.SkipCheck, "skipCheck",
		true, "Skip checking expiry for Google OAuth Token")

	RootCmd.AddCommand(apis.Cmd)
	RootCmd.AddCommand(org.Cmd)
	RootCmd.AddCommand(sync.Cmd)
	RootCmd.AddCommand(env.Cmd)
	RootCmd.AddCommand(products.Cmd)
	RootCmd.AddCommand(developers.Cmd)
	RootCmd.AddCommand(apps.Cmd)
	RootCmd.AddCommand(sharedflows.Cmd)
	RootCmd.AddCommand(kvm.Cmd)
	RootCmd.AddCommand(flowhooks.Cmd)
	RootCmd.AddCommand(targetservers.Cmd)
	RootCmd.AddCommand(token.Cmd)
	RootCmd.AddCommand(keystores.Cmd)
	RootCmd.AddCommand(keyaliases.Cmd)
	RootCmd.AddCommand(cache.Cmd)
	RootCmd.AddCommand(res.Cmd)
	RootCmd.AddCommand(projects.Cmd)
	RootCmd.AddCommand(iam.Cmd)
}

func initConfig() {
	viper.SetEnvPrefix("APIGEE")
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetConfigType("json")
}

// GetRootCmd returns the root of the cobra command-tree.
func GetRootCmd() *cobra.Command {
	return RootCmd
}
