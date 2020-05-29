package cmd

import "github.com/spf13/cobra"

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a shell.",
	Long: `Delete a shell`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	addInt(args)
	// },
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
