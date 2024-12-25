package cmd

import (
	"github.com/aqyuki/wkm/cmd/group"
	"github.com/aqyuki/wkm/config"
	"github.com/aqyuki/wkm/logging"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "root",
	Short:   "show sandboxes root",
	GroupID: group.CoreCommandGroupID,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.LoadConfig()
		if err != nil {
			logging.Errorf("failed to load config because of %v", err)
			return
		}
		logging.Plainln(config.MustAbsRoot())
	},
}
