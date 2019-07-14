package sharedflows

import (
	"../shared"
	getsf "./getsf"
	listsf "./listsf"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "sharedflows",
	Short: "Manage Apigee shared flows in an org",
	Long:  "Manage Apigee shared flows in an org",
}

func init() {

	Cmd.PersistentFlags().StringVarP(&shared.RootArgs.Org, "org", "o",
		"", "Apigee organization name")

	Cmd.MarkPersistentFlagRequired("org")
	Cmd.AddCommand(listsf.Cmd)
	Cmd.AddCommand(getsf.Cmd)
}
