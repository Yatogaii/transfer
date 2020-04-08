package lanzous

import (
	"fmt"
	"github.com/spf13/cobra"
	"transfer/apis"
)

var (
	Backend = new(lanzous)
)

type lanzous struct {
	apis.Backend
	resp     string
	Config   wssOptions
	Commands [][]string
}

func (b *lanzous) SetArgs(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&b.Config.token, "token", "t", "", "Set your user token (required)")
	cmd.Long = fmt.Sprintf("lanzous - https://www.lanzous.com/\n\n" +
		"\n  Note: This backend only supports login users. (use -t to set token)\n")
}