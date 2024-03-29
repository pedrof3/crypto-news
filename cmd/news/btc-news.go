/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package news

import (
	"fmt"

	"github.com/spf13/cobra"
)

// btcNewsCmd represents the btcNews command
var btcNews = &cobra.Command{
	Use:   "btc",
	Short: "Bitcoin news",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("btcNews called")
	},
}

// func init() {
// 	rootCmd.AddCommand(btcNewsCmd)

// 	Here you will define your flags and configuration settings.

// 	Cobra supports Persistent Flags which will work for this command
// 	and all subcommands, e.g.:
// 	btcNewsCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	Cobra supports local flags which will only run when this command
// 	is called directly, e.g.:
// 	btcNewsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
