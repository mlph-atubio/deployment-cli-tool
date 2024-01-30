/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"os"
	"gopkg.in/yaml.v2"
	"log"
	"path/filepath"
	cp "github.com/otiai10/copy"
)

type BackupConfig struct {
	FolderSource      []string `yaml:"folderSource"`
	FolderDestination string `yaml:"folderDestination"`
}

// createBackupCmd represents the create-backup command
var createBackupCmd = &cobra.Command{
	Use:   "create-backup",
	Short: "Creates a backup in a folder specified in the config file.",
	Long: `Creates a backup in a folder specified in the config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create-backup called")

		var cfg BackupConfig
		
		file, err := os.Open("C:/Users/MLMNL-ATUBIO/Documents/Projects/deployment-cli-tool/config.yml")
		if err != nil {
			log.Fatal(err)	
		}
		defer file.Close()

		fmt.Println("Config file found... Initializing")

		decoder := yaml.NewDecoder(file)
		err = decoder.Decode(&cfg)
		
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Output: %+v\n", cfg)
		fmt.Printf("Destination: %v\n", cfg.FolderDestination)

		fmt.Println("Saving to backup destination folder")
		err = createBackup(cfg.FolderSource, cfg.FolderDestination)
		if err != nil {
			log.Fatal(err)
		} 
	},
}

func createBackup(sourcePaths []string, destinationPath string) error {
	var innerFolderName = "folder1"

	// Creates new directory if it does not exist, otherwise do nothing
	err := createDirectory(destinationPath)
	if err != nil {
		fmt.Println("here")
		return err
	}

	for _, folderPath := range sourcePaths {
		_, err := os.Stat(folderPath) 	
		if err != nil {
			continue
		}
		//fmt.Printf("%v\n", folderPath)

		err = cp.Copy(folderPath, fmt.Sprintf("%s/%s", destinationPath, innerFolderName))
		if err != nil {
			return err
		}
	}

	return nil
}

func createDirectory(path string) error {
	fmt.Printf("%v\n", path)
	_, err := os.Stat(path)
	
	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		err = createDirectory(filepath.Dir(path))
		if err != nil {
			return err
		}

		err = os.Mkdir(path, 0777)
		if err != nil {
			return err
		}
	}
	
	return nil
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
