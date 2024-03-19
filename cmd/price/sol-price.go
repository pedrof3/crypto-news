/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package price

import (
	"fmt"

	"github.com/spf13/cobra"
)

// solPriceCmd represents the solPrice command
var solPrice = &cobra.Command{
	Use:   "sol",
	Short: "Solana current price in USD",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Solana current price in USD")
	},
}

// func init() {
// 	rootCmd.AddCommand(solPriceCmd)

// 	Here you will define your flags and configuration settings.

// 	Cobra supports Persistent Flags which will work for this command
// 	and all subcommands, e.g.:
// 	solPriceCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	Cobra supports local flags which will only run when this command
// 	is called directly, e.g.:
// 	solPriceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
