/*
Copyright © 2023 Bersen Naidoo bersen.naidoo@gmail.com
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bersennaidoo/devops/foundation/cert"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// Config is the internal representation of tls.yaml file.
type Config struct {
	CACert *cert.CACert          `yaml:"caCert"`
	Cert   map[string]*cert.Cert `yaml:"certs"`
}

var cfgFilePath string
var config Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "devops",
	Short: "devops is a command line tool for TLS.",
	Long: `devops is a command line tool for TLS.
	Mainly used for generation of X.509 certificates, but can be extended`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// init initializes cli application.
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFilePath, "config", "c", "", "config file (default is tls.yaml)")
}

// initConfig sets up and initializes CA config file.
func initConfig() {
	if cfgFilePath == "" {
		cfgFilePath = "tls.yaml"
	}
	cfgFileBytes, err := ioutil.ReadFile(cfgFilePath)
	if err != nil {
		fmt.Printf("Error while reading config file: %s\n", err)
		return
	}
	err = yaml.Unmarshal(cfgFileBytes, &config)
	if err != nil {
		fmt.Printf("Error while parsing config file: %s\n", err)
		return
	}
}
