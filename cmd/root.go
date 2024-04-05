/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/pedrof3/crypto-news/cmd/ath"
	"github.com/pedrof3/crypto-news/cmd/news"
	"github.com/pedrof3/crypto-news/cmd/price"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crypto-news",
	Short: "Daily cryptocurrency news",
	Long:  `CLI application for cryptocurrency news, statistics, graphics and more`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
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

func printDateTime() {
	y, m, d := time.Now().Date()
	hour, min, sec := time.Now().Clock()

	fmt.Printf("%d %s %d | %.2d:%.2d:%.2d\n", d, m, y, hour, min, sec)
}

func init() {
	rootCmd.AddCommand(ath.AthCmd)
	rootCmd.AddCommand(news.NewsCmd)
	rootCmd.AddCommand(price.PriceCmd)

	printDateTime()

}
