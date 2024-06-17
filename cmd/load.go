/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	url 			string
	totalRequests 	int
	concurrency 	int
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Perform load testing on a web service",
	Long: `Loadtester is a CLI tool that allows you to perform load testing on web services.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("load called")
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)

	loadCmd.Flags().StringVarP(&url, "url", "u", "", "The URL of the web service to be tested")
	loadCmd.Flags().IntVarP(&totalRequests, "requests", "r", 0, "The total number of requests to be made")
	loadCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 0, "The number of concurrent requests to be made")
	loadCmd.MarkFlagRequired("url")
	loadCmd.MarkFlagRequired("requests")
	loadCmd.MarkFlagRequired("concurrency")
}
