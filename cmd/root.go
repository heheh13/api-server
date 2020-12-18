package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "heheh",
	Short: "this a the main commad",
	Long: `this a  long command 
	for now we are not sure what to do with it`,
	Version: "v1.1.1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello")
		cmd.SetVersionTemplate("v2.23.3")
	},
}

//Execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
