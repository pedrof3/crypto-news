/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package news

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ethNewsCmd represents the ethNews command
var ethNews = &cobra.Command{
	Use:   "eth",
	Short: "Ethereum news",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ethNews called")
	},
}

// func init() {
// 	rootCmd.AddCommand(ethNewsCmd)

// 	Here you will define your flags and configuration settings.

// 	Cobra supports Persistent Flags which will work for this command
// 	and all subcommands, e.g.:
// 	ethNewsCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	Cobra supports local flags which will only run when this command
// 	is called directly, e.g.:
// 	ethNewsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
