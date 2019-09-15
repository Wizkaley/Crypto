package cobra

import (
	"fmt"
	"secret"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "gets secret from your secret store",
	Run: func(cmd *cobra.Command, args []string) {
		//v := se
		//fmt.Printf("key=%s\n", encodingKey)
		v := secret.File(encodingKey, secretsPath())
		fmt.Print(args)
		key := args[0]
		value, err := v.Get(key)
		//v.Get(args[0])
		if err != nil {
			fmt.Println("No value Set")
		}
		fmt.Printf("%s=%s\n", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
