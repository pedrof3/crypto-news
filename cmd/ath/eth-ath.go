/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package ath

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ethAthCmd represents the ethAth command
var ethAth = &cobra.Command{
	Use:   "eth",
	Short: "All-time highest price for Ethereum",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("All-time highest price for Ethereum")
	},
}

// func init() {
// 	rootCmd.AddCommand(ethAthCmd)

// 	Here you will define your flags and configuration settings.

// 	Cobra supports Persistent Flags which will work for this command
// 	and all subcommands, e.g.:
// 	ethAthCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	Cobra supports local flags which will only run when this command
// 	is called directly, e.g.:
// 	ethAthCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
