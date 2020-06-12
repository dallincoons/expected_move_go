/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
