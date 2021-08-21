
package root

import (
	"github.com/ondrejsika/osu/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "osu",
	Short: "osu, " + version.Version,
}
