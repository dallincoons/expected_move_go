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
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		from, _ := cmd.Flags().GetString("from")
		to, _ := cmd.Flags().GetString("to")

		PullPrices(from, to)
	},
}

func PullPrices(from string, to string) {
	pricesController := &prices.HistoricalPriceController{
		Client: http.DefaultClient,
		Tickers: prices.GetTickers(),
	}

	tz, _ := time.LoadLocation("America/Boise")

	fromDate, err := time.ParseInLocation("2006-01-02", from, tz)
	toDate, err := time.ParseInLocation("2006-01-02", to, tz)

	fmt.Println(fromDate)
	fmt.Println(toDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	pricesController.GetAllDayPricesForRange(fromDate, toDate, getWriteStrategy())
}

func getWriteStrategy() prices.DisplayStrategyInterface {
	return &prices.WritePostgres{
		Dsn: fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_DATABASE"),
		),
	}
}

func init() {
	godotenv.Load()

	rootCmd.AddCommand(syncAllCmd)

	loc, _ := time.LoadLocation("America/Boise")

	syncAllCmd.Flags().String("from", time.Now().In(loc).AddDate(0,0,0).Format("2006-01-02"), "")
	syncAllCmd.Flags().String("to", time.Now().In(loc).Format("2006-01-02"), "")
}
