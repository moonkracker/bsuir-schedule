package cmd

import (
	"github.com/withfig/autocomplete-tools/integrations/cobra"
)

func init() {
	rootCmd.AddCommand(cobracompletefig.CreateCompletionSpecCommand())
}
