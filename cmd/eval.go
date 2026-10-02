package cmd

import (
	"fmt"
	"math_engine/lexer"
	"math_engine/parser"

	"github.com/spf13/cobra"
)

func evaluate(expr string) (float64, error) {
	tokens := lexer.NewLexer(expr).Lex()
	return parser.NewParser(tokens).Parse()
}

var evalCmd = &cobra.Command{
	Use:   "eval [expression]",
	Short: "Evaluate a mathematical expression",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := evaluate(args[0])

		if err != nil {
			return err
		}

		fmt.Println(result)
		return nil
	},
}

func init() {
	rootCommand.AddCommand(evalCmd)
}
