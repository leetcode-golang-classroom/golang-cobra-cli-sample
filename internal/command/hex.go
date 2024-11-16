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
