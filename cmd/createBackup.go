/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type backupConfig struct {
	folderSource []string,
	folderDestination string,
}

// createBackupCmd represents the createBackup command
var createBackupCmd = &cobra.Command{
	Use:   "create-backup",
	Short: "Creates a backup in a folder specified in the config file.",
	Long: `Creates a backup in a folder specified in the config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create-backup called")
	},
}

func init() {
	rootCmd.AddCommand(createBackupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createBackupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createBackupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
