package secret_test

import (
	"testing"

	"github.com/rendon/testcli"
)

func TestSecret(t *testing.T) {
	testcli.Run("./cli")

	// cli, ok := os.LookupEnv("cli")
	// if !ok {
	// 	t.Errorf("..............%v", ok)
	// }
	//fmt.Println(cli)
	//testcli.
	if !testcli.Success() {
		t.Errorf("Failed to run command")
	}
	if testcli.StdoutContains("secret is An API to keep your Keys Encrypted and Safe") {
		t.Errorf("Expected %q but got %q", "secret is An API to keep your Keys Encrypted and Safe", testcli.Stdout())
	}
}

func TestCliWithGet(t *testing.T) {
	c := testcli.Command("./cli", "get", "[twitter_api]")
	c.Run()

	if !c.Success() {
		t.Errorf("Failed to run command")
	}

	if c.StdoutContains("[[twitter_api]]") {
		t.Errorf("Expected %q but got %q", "[[twitter_api]]", c.Stdout())
	}
}

func TestCliWithSet(t *testing.T) {
	c := testcli.Command("./cli", "set", "-k", "1234")

	c.Run()

	if !c.Success() {
		t.Errorf("Failed to run command")
	}

	if c.StdoutContains("key=1234\n") {
		t.Errorf("Expected %q but got %q", "key=1234\n", c.Stdout())
	}
}
