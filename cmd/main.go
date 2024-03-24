package cmd

import (
	"cobra/service"

	"github.com/spf13/cobra"
)

var cmdMain = &cobra.Command{
	Use:   "main",
	Short: "Anh print hello world",
	Run: func(cmd *cobra.Command, args []string) {
		RunService()
	},
}

func RunService() {
	var anh service.Anh
	anh.PrintHelloWorld()
}
