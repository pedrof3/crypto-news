/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package news

import (
	"fmt"

	"github.com/spf13/cobra"
)

// solNewsCmd represents the solNews command
var solNews = &cobra.Command{
	Use:   "sol",
	Short: "Solana news",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("solNews called")
	},
}

// func init() {
// 	rootCmd.AddCommand(solNewsCmd)

// 	Here you will define your flags and configuration settings.

// 	Cobra supports Persistent Flags which will work for this command
// 	and all subcommands, e.g.:
// 	solNewsCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	Cobra supports local flags which will only run when this command
// 	is called directly, e.g.:
// 	solNewsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
