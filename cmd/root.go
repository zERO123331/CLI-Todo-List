package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	Task string

	rootCmd = &cobra.Command{
		Use:   "todo-CLI",
		Short: "todo CLI",
		Long:  "A simple todo list program but for CLI",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Task, "task", "t", "", "task to work with")
}
