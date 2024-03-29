/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
  "fmt"
  "log"
  "os"
  "github.com/spf13/cobra"

  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"

)


type Config struct{
  IntervalSec int64 `mapstructure:"interval"`
  EndPoint string `mapstructure:"endpoint"`
  ReqTimeOutSec int64 `mapstructure:"timeout"`
  Command string `mapstructure:"command"`
}

var(
  cfgFile string
  configStruct Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "cobraViperFinal",
  Short: "A brief description of your application",
  Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
  // Uncomment the following line if your bare application
  // has an action associated with it:
  //	Run: func(cmd *cobra.Command, args []string) { },
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
  cobra.OnInitialize(initConfig)
  rootCmd.AddCommand(callApiCmd)
  // Here you will define your flags and configuration settings.
  // Cobra supports persistent flags, which, if defined here,
  // will be global for your application.

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobraViperFinal.yaml)")


  // Cobra also supports local flags, which will only run
  // when this action is called directly.
  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".cobraViperFinal" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".cobraViperFinal")
  }

  err := viper.ReadInConfig() // Find and read the config file
  if err != nil { // Handle errors reading the config file
    log.Fatalf("Fatal error config file: %s \n", err)
  }
  err = viper.Unmarshal(&configStruct)
  if err != nil {
    log.Fatalf("unable to decode into struct, %v", err)
  }
  log.Printf("new values are:%+v",configStruct)
}

