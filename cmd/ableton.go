/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/FileOrganize/domain"
	"github.com/spf13/cobra"
)

// abletonCmd represents the ableton command
var abletonCmd = &cobra.Command{
	Use:   "ableton",
	Short: "ableton is organized music file(mp3, wav).",
	Long: `ableton is organized music file(mp3, wav).
	
The ableton command extracts music files (mp3, wav) from the project file in the ableton work directory to the specified directory.`,
	Run: func(cmd *cobra.Command, args []string) {
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

		err := domain.Ableton(target, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(abletonCmd)
	
	abletonCmd.Flags().StringP("target", "t", "", "Target Directory.")
	abletonCmd.Flags().StringP("output", "o", "", "Output Directory.")
}
