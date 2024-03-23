/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package news

import (
	"github.com/spf13/cobra"
)

var NewsCmd = &cobra.Command{
	Use:   "news",
	Short: "Daily news in cryptocurrency world",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	NewsCmd.AddCommand(btcNews)
	NewsCmd.AddCommand(ethNews)
	NewsCmd.AddCommand(solNews)
	NewsCmd.AddCommand(defiNews)
}
