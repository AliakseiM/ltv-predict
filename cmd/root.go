package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/AliakseiM/ltv-predict/internal/flags"
	"github.com/AliakseiM/ltv-predict/internal/models"
)

var (
	model     string
	source    string
	aggregate string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ltv-predict",
	Short: "A brief description",
	Long:  `A longer description`,
	RunE: func(_ *cobra.Command, _ []string) error {
		if err := flags.ValidateValues(model, source, aggregate); err != nil {
			return err
		}

		switch models.SourceType(source) {
		case models.SourceTypeJSON:
			err := models.PrintJSONInput()
			if err != nil {
				return err
			}
		case models.SourceTypeCSV:
			err := models.PrintCSVInput()
			if err != nil {
				return err
			}
		default:
			return nil
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&model, flags.Model.String(), flags.Model.Shorthand(), "", "Model")
	if err := rootCmd.MarkFlagRequired(flags.Model.String()); err != nil {
		cobra.CheckErr(err)
	}

	rootCmd.Flags().StringVarP(&source, flags.Source.String(), flags.Source.Shorthand(), "", "CSV or JSON data")
	if err := rootCmd.MarkFlagRequired(flags.Source.String()); err != nil {
		cobra.CheckErr(err)
	}

	rootCmd.Flags().StringVarP(&aggregate, flags.Aggregate.String(), flags.Aggregate.Shorthand(), "", "grouping by country or campaign")
	if err := rootCmd.MarkFlagRequired(flags.Aggregate.String()); err != nil {
		cobra.CheckErr(err)
	}
}
