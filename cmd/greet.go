package cmd

import (
	"fmt"

	"github.com/ahmadu2305/go-demo-1/mascot"
	"github.com/spf13/cobra"
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Print the best mascot",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(mascot.BestMascot())
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)
}
