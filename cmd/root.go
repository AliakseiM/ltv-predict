package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	flagModel     = "model"
	flagSource    = "source"
	flagAggregate = "aggregate"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(model)
		fmt.Println(source)
		fmt.Println(aggregate)
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
	rootCmd.Flags().StringVarP(&model, flagModel, "m", "", "Model")
	if err := rootCmd.MarkFlagRequired(flagModel); err != nil {
		cobra.CheckErr(err)
	}

	rootCmd.Flags().StringVarP(&source, flagSource, "s", "", "CSV or JSON data")
	if err := rootCmd.MarkFlagRequired(flagSource); err != nil {
		cobra.CheckErr(err)
	}

	rootCmd.Flags().StringVarP(&aggregate, flagAggregate, "a", "", "grouping by country or campaign")
	if err := rootCmd.MarkFlagRequired(flagAggregate); err != nil {
		cobra.CheckErr(err)
	}
}
