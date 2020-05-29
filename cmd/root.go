package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker-shell",
	Short: "docker-shell is a script used to create isolated shells using Docker.",
	Long:  `docker-shell is a script used to create isolated shells using Docker.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
