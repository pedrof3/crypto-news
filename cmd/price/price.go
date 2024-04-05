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
