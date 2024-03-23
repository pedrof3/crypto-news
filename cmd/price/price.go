/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package price

import (
	"github.com/spf13/cobra"
)

// priceCmd represents the price command
var PriceCmd = &cobra.Command{
	Use:   "price",
	Short: "Tokens price.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	PriceCmd.AddCommand(btcPrice)
	PriceCmd.AddCommand(ethPrice)
	PriceCmd.AddCommand(solPrice)
	PriceCmd.AddCommand(top10Price)
	PriceCmd.AddCommand(top20Price)
}

// func init() {
// 	rootCmd.AddCommand(priceCmd)

// 	Here you will define your flags and configuration settings.

// 	Cobra supports Persistent Flags which will work for this command
// 	and all subcommands, e.g.:
// 	priceCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	Cobra supports local flags which will only run when this command
// 	is called directly, e.g.:
// 	priceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
