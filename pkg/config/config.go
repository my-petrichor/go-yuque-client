package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ConfigFile string
)

func Init() {
	if ConfigFile != "" {
		viper.SetConfigFile(ConfigFile)
	} else {
		homeDir, err := os.UserHomeDir()
		cobra.CheckErr(err)

		if err := checkFileExist(filepath.Join(homeDir, ".yuque.yaml")); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(homeDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".yuque")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Fatal(err)
	}
}

func checkFileExist(file string) error {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create(file)
			if err != nil {
				return err
			}
			defer f.Close()
		} else {
			return err
		}
	}

	return nil
}
