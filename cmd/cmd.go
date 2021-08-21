package cmd

import (
	"github.com/ondrejsika/osu/cmd/root"
	_ "github.com/ondrejsika/osu/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
