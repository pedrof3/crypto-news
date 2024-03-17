/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// priceCmd represents the price command
var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "Tokens price. Show top 10 coins by default",
	Long: `PRICE
	crypto-news --price btc: show bitcoin price
	crypto-news --price 20: list price for top 20 coins by market cap`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("price called")
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// priceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// priceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
