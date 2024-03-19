/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package news

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newsCmd represents the news command
var NewsCmd = &cobra.Command{
	Use:   "news",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("news called")
	},
}

// func init() {
// 	rootCmd.AddCommand(newsCmd)

// 	Here you will define your flags and configuration settings.

// 	Cobra supports Persistent Flags which will work for this command
// 	and all subcommands, e.g.:
// 	newsCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	Cobra supports local flags which will only run when this command
// 	is called directly, e.g.:
// 	newsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
