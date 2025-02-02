package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd đại diện cho lệnh gốc
var rootCmd = &cobra.Command{
	Use:   "crud-generator",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from rootCmd!")
	},
}

// Execute thực thi lệnh gốc
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
