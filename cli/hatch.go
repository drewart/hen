package cli

import (
	"fmt"
	"os"

	"github.com/drewart/hen"
	"github.com/spf13/cobra"
)

var hatchCmd = &cobra.Command{
	Use:   "hatch",
	Short: "hatch hen eggs",
	Long:  `hatch hen eggs`,
	Run: func(cmd *cobra.Command, args []string) {
		f, _ := cmd.Flags().GetString("file")
		hatch(f)
	},
}

func init() {
	rootCmd.AddCommand(hatchCmd)
	hatchCmd.Flags().StringP("file", "f", "", "file to parse for egg()")
}

func hatch(filePath string) {
	err := hen.Hatch(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
