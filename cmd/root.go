package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"

	"github.com/AliakseiM/ltv-predict/internal/datasource/csv"
	"github.com/AliakseiM/ltv-predict/internal/datasource/json"
	"github.com/AliakseiM/ltv-predict/internal/flags"
	"github.com/AliakseiM/ltv-predict/internal/models"
	"github.com/AliakseiM/ltv-predict/internal/predictor/linearRegression"
)

var (
	model     string
	source    string
	aggregate string
)

const (
	day = 60
)

const (
	jsonFile = "data/test_data.json"
	csvFile  = "data/test_data.csv"
)

type Datasource interface {
	LoadData() error
	GroupBy(models.AggregateType)
	Prepare() (map[string][]float64, error)
	Print()
}

type Predictor interface {
	PredictForDay(data []float64, day int) (float64, error)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ltv-predict",
	Short: "A brief description",
	Long:  `A longer description`,
	RunE: func(_ *cobra.Command, _ []string) error {
		if err := flags.ValidateValues(model, source, aggregate); err != nil {
			return err
		}

		var ds Datasource

		switch models.SourceType(source) {
		case models.SourceTypeJSON:
			ds = json.NewDatasource(jsonFile)
		case models.SourceTypeCSV:
			ds = csv.NewDatasource(csvFile)
		default:
			// TODO: return error
			return nil
		}

		if err := ds.LoadData(); err != nil {
			return err
		}

		ds.GroupBy(models.AggregateType(aggregate))

		prepared, err := ds.Prepare()
		if err != nil {
			return err
		}

		// TODO: predict

		var predictor Predictor
		switch models.PredictionModel(model) {
		case models.LinearRegression:
			predictor = linearRegression.New()
		case models.ExponentialSmoothing:
			// TODO
		default:
			// TODO: return error
			return nil
		}

		gr := new(errgroup.Group)

		for group, data := range prepared {
			group, data := group, data
			gr.Go(func() error {
				predicted, err := predictor.PredictForDay(data, 60)
				if err != nil {
					return err
				}

				fmt.Printf("%s: %.2f\n", group, predicted)

				return nil
			})
		}

		if err := gr.Wait(); err != nil {
			return err
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
