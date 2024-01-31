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
	BackupMap         []struct{
		Source        string `yaml:"source"` 
		Destination   string `yaml:"destination"` 
	} `yaml:"backupMap"`
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

		log.Println("Config file found... Initializing")

		decoder := yaml.NewDecoder(file)
		err = decoder.Decode(&cfg)
		
		if err != nil {
			log.Fatal(err)
		}

		for _, folderPair := range cfg.BackupMap {	
			err = createBackup(folderPair.Source, folderPair.Destination)
			if err != nil {
				log.Fatal(err)
			}
		}

		log.Println("Backup successfully created.")
	},
}

func createBackup(sourcePath string, destinationPath string) error {
	var innerFolderName = "folder1"
	fmt.Printf("Source path: %s\n", sourcePath)

	// Creates new directory if it does not exist, otherwise do nothing
	err := createDirectory(destinationPath)
	if err != nil {
		log.Printf("Error creating directory: %v\n", destinationPath)
		return err
	}

	_, err = os.Stat(sourcePath)	
	if err != nil {
		return err	
	}

	log.Printf("Creating backup files...")
	err = cp.Copy(sourcePath, fmt.Sprintf("%s/%s", destinationPath, innerFolderName))
	if err != nil {
		return err
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

		log.Printf("Directory %v does not exist. Creating directory...\n", path)
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
