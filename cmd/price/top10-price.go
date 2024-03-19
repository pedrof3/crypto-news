/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package price

import (
	"fmt"

	"github.com/spf13/cobra"
)

// top10PriceCmd represents the top10Price command
var top10Price = &cobra.Command{
	Use:   "top10",
	Short: "Current price in USD for top 10 coins by market cap",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current price in USD for top 10 coins by market cap")
	},
}

// func init() {
// 	rootCmd.AddCommand(top10PriceCmd)

// 	Here you will define your flags and configuration settings.

// 	Cobra supports Persistent Flags which will work for this command
// 	and all subcommands, e.g.:
// 	top10PriceCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	Cobra supports local flags which will only run when this command
// 	is called directly, e.g.:
// 	top10PriceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
