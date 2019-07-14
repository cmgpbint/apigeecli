package getprod

import (
	"../../shared"
	"github.com/spf13/cobra"
	"net/url"
	"path"
)

var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Gets an API product from an organization",
	Long:  "Gets an API product from an organization",
	Run: func(cmd *cobra.Command, args []string) {
		u, _ := url.Parse(shared.BaseURL)
		u.Path = path.Join(u.Path, shared.RootArgs.Org, "apiproducts", name)
		shared.GetHttpClient(u.String(), shared.RootArgs.Token)
	},
}

var name string

func init() {

	Cmd.Flags().StringVarP(&shared.RootArgs.Org, "org", "o",
		"", "Apigee organization name")

	Cmd.Flags().StringVarP(&name, "name", "n",
		"", "name of the API Product")

	Cmd.MarkFlagRequired("org")
	Cmd.MarkFlagRequired("name")
}
