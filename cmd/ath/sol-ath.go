/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package ath

import (
	"fmt"

	"github.com/spf13/cobra"
)

// solAthCmd represents the solAth command
var solAth = &cobra.Command{
	Use:   "sol",
	Short: "All-time highest price for Solana",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("All-time highest price for Solana")
	},
}

// func init() {
// 	rootCmd.AddCommand(solAthCmd)

// 	Here you will define your flags and configuration settings.

// 	Cobra supports Persistent Flags which will work for this command
// 	and all subcommands, e.g.:
// 	solAthCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	Cobra supports local flags which will only run when this command
// 	is called directly, e.g.:
// 	solAthCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
