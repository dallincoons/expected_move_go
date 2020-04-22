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
	"log"
	"os"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Retrieve price for stock ticker",
	Long: ``,
	Run: func(cmd *
	cobra.Command, args []string) {
		ticker, _ := cmd.Flags().GetString("ticker")
		date, _ := cmd.Flags().GetString("date")
		fileName, _ := cmd.Flags().GetString("file")

		openFile := getCsvWriter(&fileName)

		pricesController := &prices.HistoricalPriceController{
			Ticker: ticker,
			Date:   date,
			WriteStrategy: &prices.WriteCSV{
				Writer: openFile,
				DisplayToUser: os.Stdout,
			},
		}

		pricesController.GetPrices()
	},
}

func getCsvWriter(fileName *string) *os.File {

	file, err := os.OpenFile(*fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		log.Fatalf("failed opening file, %s", err)
	}

	return file
}

func init() {
	rootCmd.AddCommand(syncCmd)

	syncCmd.Flags().String("ticker", "SPY", "Ticker to retrieve price for")
	syncCmd.Flags().String("date", "", "Date to pull historical price for")
	syncCmd.Flags().String("file", "", "file to write price to")
}
