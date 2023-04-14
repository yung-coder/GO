package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var existsCmd = &cobra.Command{
	Use:   "present",
	Short: "Trace the IP",
	Long:  `Trace the IP.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, name := range args {
			_, err := os.Stat(name)

			if err != nil {
				fmt.Println("File dose not exists")
			} else {
				fmt.Println("Here")

			}

		}
	},
}

func init() {
	rootCmd.AddCommand(existsCmd)
}
