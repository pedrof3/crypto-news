/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	btc   string
	eth   string
	sol   string
	top10 string
	top20 string
	top30 string
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
	athCmd.Flags().StringVarP(&btc, "bitcoin", "b", "", "bitcoin all time high")
	athCmd.Flags().StringVarP(&eth, "ethereum", "e", "", "ethereum all time high")
	athCmd.Flags().StringVarP(&sol, "solana", "s", "", "solana all time high")
	athCmd.Flags().StringVarP(&top10, "top10", "1", "", "all time high for top 10 coins")
	athCmd.Flags().StringVarP(&top20, "top20", "2", "", "all time high for top 20 coins")
	athCmd.Flags().StringVarP(&top30, "top30", "3", "", "all time high for top 30 coins")

	rootCmd.AddCommand(athCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// athCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// athCmd.Flags().BoolVarP("toggle", "t", true, "Help message for toggle")
}