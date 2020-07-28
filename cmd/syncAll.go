package cmd

import (
	"expected_move/prices"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"time"
	"github.com/joho/godotenv"
)

var syncAllCmd = &cobra.Command{
	Use:   "syncAll",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		date, _ := cmd.Flags().GetString("date")
		//filepath, _ := cmd.Flags().GetString("filepath")

		pricesController := &prices.HistoricalPriceController{
			Client: http.DefaultClient,
		}

		pricesController.GetPrices(date, prices.GetTickers(), getWriteStrategy())
	},
}

func createWriter(filepath string) map[string]prices.DisplayStrategyInterface {
	writers := make(map[string]prices.DisplayStrategyInterface)
	for _, ticker := range prices.GetTickers() {
		writers[ticker] = prices.GetWriteStrategy(fmt.Sprintf("%s/%s.csv", filepath, ticker))
	}

	return writers
}

func getWriteStrategy() prices.DisplayStrategyInterface {
	return &prices.WritePostgres{
		Dsn: fmt.Sprintf("postgresql://%s:%s@localhost:5432/%s?sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DATABASE"),
		),
	}
}

func init() {
	godotenv.Load()

	rootCmd.AddCommand(syncAllCmd)

	syncAllCmd.Flags().String("date", time.Now().Format("2006-01-02"), "Date to pull historical prices for")
}
