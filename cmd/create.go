/*
Copyright © 2023 Bersen Naidoo bersen.naidoo@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create CA, certs, or keys",
	Long:  `commands to create resources (ca, certs, keys)`,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
