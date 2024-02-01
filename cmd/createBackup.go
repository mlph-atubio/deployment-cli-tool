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
	"errors"
)

type BackupConfig struct {
	BackupMap         []struct{
		Source        string `yaml:"source"` 
		Destination   string `yaml:"destination"` 
	} `yaml:"backupMap"`
}

var (
	ConfigSource string
	FolderName string
)

// createBackupCmd represents the create-backup command
var createBackupCmd = &cobra.Command{
	Use:   "create-backup",
	Short: "Creates a backup in a folder specified in the config file.",
	Long: `Creates a backup in a folder specified in the config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg BackupConfig

		if ConfigSource == "" {
			log.Println("Config file not set in the flag. Checking current directory.")
			ConfigSource = "./config.yml"
		}
		
		file, err := os.Open(ConfigSource)
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
	if FolderName == "" {
		errors.New("Empty folder name")	
	}

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

	basePath := filepath.Base(sourcePath)
	log.Printf("Creating backup for %s.\n", basePath)
	err = cp.Copy(sourcePath, fmt.Sprintf("%s/%s/%s", destinationPath, FolderName, basePath))
	if err != nil {
		return err
	}

	return nil
}

func createDirectory(path string) error {
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
	createBackupCmd.Flags().StringVarP(&ConfigSource, "configSource", "s", "", "File location of configuration file")
	createBackupCmd.Flags().StringVarP(&FolderName, "folderName", "f", "", "Folder name for the backup")
	rootCmd.AddCommand(createBackupCmd)
}
