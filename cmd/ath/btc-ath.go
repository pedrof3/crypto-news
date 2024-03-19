/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package ath

import (
	"fmt"

	"github.com/spf13/cobra"
)

// btcAthCmd represents the btcAth command
var btcAth = &cobra.Command{
	Use:   "btc",
	Short: "All-time highest price for Bitcoin",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("All-time highest price for Bitcoin")
	},
}

func init() {
	// rootCmd.AddCommand(btcAthCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// btcAthCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// btcAthCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
