package cobra

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "secret",
	Short: "secret is An API to keep your Keys Encrypted and Safe",
}

var encodingKey string

func init() {
	RootCmd.PersistentFlags().StringVarP(&encodingKey, "key", "k", "", "the ke to use while encoing and decoding")
}
func secretsPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".secrets")
	//return home
}
