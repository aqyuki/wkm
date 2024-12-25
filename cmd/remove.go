package cmd

import (
	"github.com/aqyuki/wkm/cmd/group"
	"github.com/aqyuki/wkm/config"
	"github.com/aqyuki/wkm/logging"
	"github.com/aqyuki/wkm/sandbox"
	"github.com/spf13/cobra"
)

var RemoveCmd = &cobra.Command{
	Aliases: []string{"rm"},
	Use:     "remove",
	Short:   "Remove a sandbox directory",
	GroupID: group.CoreCommandGroupID,
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		config, err := config.LoadConfig()
		if err != nil {
			logging.Errorf("failed to load config because of %v", err)
			return
		}

		if err := sandbox.Remove(config.Root, title); err != nil {
			logging.Errorf("failed to remove sandbox directory because of %v", err)
			return
		}

		logging.Infof("sandbox, %s was removed\n", title)
	},
}
