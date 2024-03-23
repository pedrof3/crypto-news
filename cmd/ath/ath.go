/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package ath

import (
	"github.com/spf13/cobra"
)

// athCmd represents the ath command
var AthCmd = &cobra.Command{
	Use:   "ath",
	Short: "All time highest price",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	AthCmd.AddCommand(btcAth)
	AthCmd.AddCommand(ethAth)
	AthCmd.AddCommand(solAth)
	AthCmd.AddCommand(top10Ath)
	AthCmd.AddCommand(top20Ath)
}

// func init() {

// 	rootCmd.AddCommand(athCmd)

// 	Here you will define your flags and configuration settings.

// 	Cobra supports Persistent Flags which will work for this command
// 	and all subcommands, e.g.:
// 	athCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	Cobra supports local flags which will only run when this command
// 	is called directly, e.g.:
// 	athCmd.Flags().BoolVarP("toggle", "t", true, "Help message for toggle")
// }
