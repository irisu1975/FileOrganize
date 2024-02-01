/*
Copyright © 2024 OrganizeFile <galleta.138@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/FileOrganize/domain"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "FileOrganize",
	Short: "FileOrganize is organized file.",
	Long: `FileOrganize is organized file.
	
You can organize file by creating directories for each extension.`,
	Run: func(cmd *cobra.Command, args []string){
		target, _ := cmd.Flags().GetString("target")
		if target == "" {
			fmt.Println("Please specify -t(--target)")
			return
		}

		output, _ := cmd.Flags().GetString("output")
		if output == "" {
			fmt.Println("Please specify -o(--output)")
			return
		}

		config, _ := cmd.Flags().GetString("config")

		err := domain.OrganizeFile(target, output, config)
		if err != nil {
			fmt.Println(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("target", "t", "", "Target Directory.")
	rootCmd.Flags().StringP("output", "o", "", "Output Directory.")
	rootCmd.Flags().StringP("config", "c", "", "Specify config json file.")
}