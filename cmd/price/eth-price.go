/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package price

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ethPriceCmd represents the ethPrice command
var ethPrice = &cobra.Command{
	Use:   "eth",
	Short: "Ethereum current price in USD",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ethereum current price in USD")
	},
}

// func init() {
// 	rootCmd.AddCommand(ethPriceCmd)

// 	Here you will define your flags and configuration settings.

// 	Cobra supports Persistent Flags which will work for this command
// 	and all subcommands, e.g.:
// 	ethPriceCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	Cobra supports local flags which will only run when this command
// 	is called directly, e.g.:
// 	ethPriceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
