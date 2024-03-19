/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package news

import (
	"fmt"

	"github.com/spf13/cobra"
)

// defiNewsCmd represents the defiNews command
var defiNews = &cobra.Command{
	Use:   "defi",
	Short: "Defi news",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("defiNews called")
	},
}

// func init() {
// 	rootCmd.AddCommand(defiNewsCmd)

// 	Here you will define your flags and configuration settings.

// 	Cobra supports Persistent Flags which will work for this command
// 	and all subcommands, e.g.:
// 	defiNewsCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	Cobra supports local flags which will only run when this command
// 	is called directly, e.g.:
// 	defiNewsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
