/*
Copyright © 2023 Bersen Naidoo bersen.naidoo@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/bersennaidoo/devops/foundation/key"
	"github.com/spf13/cobra"
)

var keyOut string
var keyLength int

// keyCreateCmd represents the keyCreate command
var keyCreateCmd = &cobra.Command{
	Use:   "key",
	Short: "key commands",
	Long:  `commands to create keys`,

	Run: func(cmd *cobra.Command, args []string) {
		err := key.CreateRSAPrivateKeyAndSave(keyOut, keyLength)
		if err != nil {
			fmt.Printf("Create key error: %s\n", err)
			return
		}
		fmt.Printf("Key created %s with length %d\n", keyOut, keyLength)
	},
}

func init() {
	createCmd.AddCommand(keyCreateCmd)
	keyCreateCmd.Flags().StringVarP(&keyOut, "key-out", "k", "key.pem", "destination path for key")
	keyCreateCmd.Flags().IntVarP(&keyLength, "key-length", "l", 4096, "key length")
}
