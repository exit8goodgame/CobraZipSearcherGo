/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/exit8goodgame/CobraZipSearcherGo/handler"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zipPassSearch",
	Short: "Run zip pass search",
	Long:  `Run zip pass search`,
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := zap.NewDevelopment()
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			logger.Error("File Error.", zap.String("msg", err.Error()))
		}
		err = validateFile(file)
		if err != nil {
			logger.Error("File Error.", zap.String("msg", err.Error()))
		}
		target, _ := cmd.Flags().GetString("target")
		char, _ := cmd.Flags().GetInt("char")
		length, _ := cmd.Flags().GetInt("length")
		handler.RootHandler(file, target, char, length, logger)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("file", "f", "", "Zip File (required)")
	rootCmd.Flags().StringP("target", "t", "", "Pass Character Target (optional)")
	rootCmd.Flags().IntP("char", "c", 0, "Pass Character Type (optional): 1=short, 2=long")
	rootCmd.Flags().IntP("length", "l", 0, "Pass Search Max Length (optional)")
}

func validateFile(file string) error {
	if filepath.Ext(file) == ".zip" {
		fmt.Errorf("File extension is not zip")
	}
	return nil
}
