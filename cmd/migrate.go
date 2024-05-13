package cmd

import (
	"blog/pkg/bootstrap"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Table Migrate",
	Long:  "Table Migrate",
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

func migrate() {
	bootstrap.Migrate()
}
