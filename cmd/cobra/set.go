package cobra

import (
	"fmt"
	"secret"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets your secret in your secret store",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("key=%s\n", encodingKey)
		v := secret.File(encodingKey, secretsPath())
		key, value := args[0], args[1]
		err := v.Set(key, value)
		if err != nil {
			panic(err)
		}
		fmt.Println("Value Set Successfully")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
