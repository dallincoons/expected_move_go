package cmd

import (
	"expected_move/prices"
	"github.com/spf13/cobra"
	"net/http"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Retrieve price for stock ticker",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		ticker, _ := cmd.Flags().GetString("ticker")
		date, _ := cmd.Flags().GetString("date")
		fileName, _ := cmd.Flags().GetString("file")

		pricesController := &prices.HistoricalPriceController{
			Client: &http.Client{},
		}

		pricesController.GetPrice(ticker, date, prices.GetWriteStrategy(fileName))
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	syncCmd.Flags().String("ticker", "SPY", "Ticker to retrieve price for")
	syncCmd.Flags().String("date", "", "Date to pull historical price for")
	syncCmd.Flags().String("file", "", "file to write price to")
}
