/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// athCmd represents the ath command
var athCmd = &cobra.Command{
	Use:   "ath",
	Short: "All time high. Show to 10 coins by default",
	Long: `ALL TIME HIGH
	crypto-news --ath btc: show bitcoin all time highest price
	crypto-news --ath 20: list all time highest prices for top 20 coins by market cap`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ath called")
	},
}

func init() {
	rootCmd.AddCommand(athCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// athCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// athCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
