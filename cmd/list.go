package cmd

import (
	"github.com/aqyuki/wkm/cmd/group"
	"github.com/aqyuki/wkm/config"
	"github.com/aqyuki/wkm/logging"
	"github.com/aqyuki/wkm/sandbox"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Aliases: []string{"ls"},
	Use:     "list",
	Short:   "List all sandbox directories",
	GroupID: group.CoreCommandGroupID,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.LoadConfig()
		if err != nil {
			logging.Errorf("failed to load configuration because of %v", err)
			return
		}

		directories, err := sandbox.List(config.Root)
		if err != nil {
			logging.Errorf("failed to list sandbox directories because of %v", err)
			return
		}
		displayDirectories(directories)
	},
}

func displayDirectories(directories []string) {
	for _, dir := range directories {
		logging.Plainln(dir)
	}
}
