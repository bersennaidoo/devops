/*
Copyright © 2023 Bersen Naidoo bersen.naidoo@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/bersennaidoo/devops/foundation/cert"
	"github.com/spf13/cobra"
)

var caKey string
var caCert string

// caCreateCmd represents the caCreate command
var caCreateCmd = &cobra.Command{
	Use:   "ca",
	Short: "ca commands",
	Long:  `commands to create the CA`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cert.CreateCACert(config.CACert, caKey, caCert)
		if err != nil {
			fmt.Printf("Create CA error: %s\n", err)
			return
		}
		fmt.Printf("CA created. Key: %s, cert: %s\n", caKey, caCert)
	},
}

func init() {
	createCmd.AddCommand(caCreateCmd)

	caCreateCmd.Flags().StringVarP(&caKey, "key-out", "k", "ca.key", "destination path for ca key")
	caCreateCmd.Flags().StringVarP(&caCert, "cert-out", "o", "ca.crt", "destination path for ca cert")
}
