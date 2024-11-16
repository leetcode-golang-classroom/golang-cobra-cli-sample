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
