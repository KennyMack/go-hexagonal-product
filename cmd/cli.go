/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/kennymack/go-hexagonal-product/adapters/cli"
	"github.com/spf13/cobra"
)

var action string
var productID string
var productName string
var productPrice float64

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := cli.Run(&productService, action, productID, productName, productPrice)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	cliCmd.Flags().StringVarP(&action, "action", "a", "enable", "ENABLE/DISABLE a product")
	cliCmd.Flags().StringVarP(&productID, "id", "i", "", "Return product ID")
	cliCmd.Flags().StringVarP(&productName, "product", "n", "", "Return product Name")
	cliCmd.Flags().Float64VarP(&productPrice, "price", "p", 0.0, "Return product Price")


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
