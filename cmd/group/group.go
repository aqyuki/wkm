package group

import "github.com/spf13/cobra"

const (
	CoreCommandGroupID       = "core"
	AdditionalCommandGroupID = "additional"
)

var (
	CoreCommandGroup = &cobra.Group{
		ID:    CoreCommandGroupID,
		Title: "Core Commands",
	}

	AdditionalCommandGroup = &cobra.Group{
		ID:    AdditionalCommandGroupID,
		Title: "Additional Commands",
	}
)
