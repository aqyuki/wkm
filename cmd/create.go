package cmd

import (
	"github.com/aqyuki/wkm/cmd/group"
	"github.com/aqyuki/wkm/config"
	"github.com/aqyuki/wkm/logging"
	"github.com/aqyuki/wkm/sandbox"
	"github.com/spf13/cobra"
)

var CreateCommand = &cobra.Command{
	Use:     "create",
	Short:   "Create a new sandbox directory",
	GroupID: group.CoreCommandGroupID,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		config, err := config.LoadConfig()
		if err != nil {
			logging.Errorf("failed to load wkm config because of %v\n", err)
			return
		}

		path, err := sandbox.Create(config.Root, title)
		if err != nil {
			logging.Errorf("failed to create a new sandbox because of %v\n", err)
			return
		}

		logging.Infof("new sandbox, %s was created successfully\npath : %s\n", title, path)
	},
}
