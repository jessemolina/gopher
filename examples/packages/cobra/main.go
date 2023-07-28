package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	username string
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "An example app with Cobra",
		Long:  "An example app with Cobra and Viper",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("hello, %s\n", username)
		},
	}

	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "the username flag")

	rootCmd.AddCommand(&cobra.Command{
		Use: "command1",
		Short: "An example command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("this is command command1")
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use: "command2",
		Short: "An example command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("this is command command2")
		},
	})

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}

}
