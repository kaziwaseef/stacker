package cmd

import (
	"context"

	"github.com/kaziwaseef/stacker/internal/page"
	"github.com/kaziwaseef/stacker/internal/types"
	"github.com/spf13/cobra"
)

// doctorCmd represents the doctor command
var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Checks and reports if everything is working fine or not.",
	Long: `Checks and reports if everything is working fine or not.
it checks for dependencies and other necessary things to work properly.`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose := cmd.Flag("verbose")
		ctx := context.WithValue(cmd.Context(), types.Verbose, verbose.Value.String() == "true")
		page.DoctorPage(ctx)
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
