package cmd

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var replCommand = &cobra.Command{
	Use:   "repl",
	Short: "Start an interactive session",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		scanner := bufio.NewScanner(cmd.InOrStdin())

		fmt.Println("calc repl — type 'quit' to leave")

		for {
			fmt.Print("> ")

			if !scanner.Scan() {
				fmt.Println()
				break
			}

			line := strings.TrimSpace(scanner.Text())

			switch line {
			case "":
				continue
			case "quit":
				return nil
			}

			result, err := evaluate(line)

			if err != nil {
				fmt.Println("error:", err)
				continue
			}

			fmt.Println(result)
		}

		return scanner.Err()
	},
}

func init() {
	rootCommand.AddCommand(replCommand)
}
