/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package ath

import (
	"fmt"

	"github.com/spf13/cobra"
)

// top20AthCmd represents the top20Ath command
var top20Ath = &cobra.Command{
	Use:   "top20",
	Short: "All-time highest price for actual top 20 coins by market cap",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("All-time highest price for actual top 20 coins by market cap")
	},
}

// func init() {
// 	rootCmd.AddCommand(top20AthCmd)

// 	Here you will define your flags and configuration settings.

// 	Cobra supports Persistent Flags which will work for this command
// 	and all subcommands, e.g.:
// 	top20AthCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	Cobra supports local flags which will only run when this command
// 	is called directly, e.g.:
// 	top20AthCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
