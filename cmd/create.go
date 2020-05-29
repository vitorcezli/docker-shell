package cmd

import "github.com/spf13/cobra"

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a shell.",
	Long:  `Create a shell`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	addInt(args)
	// },
}

func init() {
	rootCmd.AddCommand(createCmd)
}
