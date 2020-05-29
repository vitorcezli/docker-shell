package cmd

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List shells.",
	Long:  `List shells`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	addInt(args)
	// },
}

func init() {
	rootCmd.AddCommand(listCmd)
}
