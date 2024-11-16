package command

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// uuidcmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generates UUID",
	Long: `Generates UUID. For example:

securerandom uuid`,
	Run: func(cmd *cobra.Command, args []string) {
		newUUID := uuid.New()
		fmt.Println(newUUID.String())
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
}
