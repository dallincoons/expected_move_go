package cmd

import (
	"expected_move/prices"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"time"
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
		from, _ := cmd.Flags().GetString("from")
		to, _ := cmd.Flags().GetString("to")

		pricesController := &prices.HistoricalPriceController{
			Client: http.DefaultClient,
		}

		tz, _ := time.LoadLocation("Local")

		fromDate, err := time.ParseInLocation("2006-01-02", from, tz)
		toDate, err := time.ParseInLocation("2006-01-02", to, tz)

		if err != nil {
			fmt.Println(err)
			return
		}

		pricesController.GetAllDayPricesForRange(fromDate, toDate, getWriteStrategy())
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
	syncAllCmd.Flags().String("from", time.Now().AddDate(0,0,0).Format("2006-01-02"), "")
	syncAllCmd.Flags().String("to", time.Now().Format("2006-01-02"), "")
}
