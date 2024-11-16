# golang-cobra-cli-sample

This repository is how to use cobra framework to write cli application

## dependency

1. cobra

```shell
go get github.com/spf13/cobra@latest
```

## logic

basic struct

main contain command.Execute method
```golang
package main

import (
	"log"

	"github.com/leetcode-golang-classroom/golang-cobra-cli-sample/internal/command"
)

func main() {
	err := command.Execute()
	if err != nil {
		log.Fatalf("command.Execute error: %v", err)
	}
}

```

root cmd for handle main Execute
```golang
import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "securerandom",
	Short: "Secure random number generators",
	Long:  "an interface to secure random number generators",
}

func Execute() error {
	return rootCmd.Execute()
}

```
hex.go for generate a random number with hex encoded
```golang
package command

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

// hexCmd
var hexCmd = &cobra.Command{
	Use:   "hex",
	Short: "Generates Hex numbers",
	Long: `Provided length, it generates hex numbers.
For example:
securerandom hex -l 10`,
	Run: func(cmd *cobra.Command, args []string) {
		length, _ := cmd.Flags().GetInt("length")
		hex, err := generateHexString(length)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(hex)
		}
	},
}
// init for add cmd, and flag setup
func init() {
	rootCmd.AddCommand(hexCmd)
	hexCmd.Flags().IntP("length", "l", 4, "Length of hex")
}

func generateHexString(length int) (string, error) {
	// Calculate the nubmer of bytes needed
	byteLength := (length + 1) / 2
	bytes := make([]byte, byteLength)
	// Generate random byte
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	// Conver to hex and truncate to the desired length
	hexString := hex.EncodeToString(bytes)[:length]
	return hexString, nil
}

```

base64.go generate random number with base64 encoded
```golang
package command

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Generate base64 string",
	Long: `Provided length, it generates base64 string.
For example:
securerandom base64 -l 10`,
	Run: func(cmd *cobra.Command, args []string) {
		length, _ := cmd.Flags().GetInt("length")
		str, err := generateBase64tring(length)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(str)
		}
	},
}

func generateBase64tring(byteLength int) (string, error) {
	// Generate random bytes
	bytes := make([]byte, byteLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	// Encode the bytes ot a Base64 string
	base64String := base64.StdEncoding.EncodeToString(bytes)
	return base64String, nil
}

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.Flags().IntP("length", "l", 4, "Length of base64")
}

```

uuid.go for generate a uuid 
```golang
package command

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// uuidcmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generates UUID",
	Long: `Generates UUID. For example:

securerandom uuid`,
	Run: func(cmd *cobra.Command, args []string) {
		newUUID := uuid.New()
		fmt.Println(newUUID.String())
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
}

```