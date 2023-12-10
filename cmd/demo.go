/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/exit8goodgame/CobraZipSearcherGo/handler"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// demoCmd represents the demo command
var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "Run demo zip pass search",
	Long:  `Run demo zip pass search.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := zap.NewDevelopment()
		handler.DemoHandler(logger)
	},
}

func init() {
	rootCmd.AddCommand(demoCmd)
}
