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
