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

		pricesController := &prices.HistoricalPriceController{
			Ticker: ticker,
			Date:   date,
			WriteStrategy: getWriteStrategy(cmd),
		}

		pricesController.GetPrices()
	},
}

func getWriteStrategy(cmd *cobra.Command) prices.DisplayStrategyInterface {

	fileName, _ := cmd.Flags().GetString("file")

	if (fileName != "") {
		return &prices.WriteCSV{
			Writer:        getCsvWriter(&fileName),
			DisplayToUser: os.Stdout,
		}
	}

	return &prices.DisplayTable{}
}

func getCsvWriter(fileName *string) *os.File {
	file, err := os.OpenFile(*fileName, os.O_APPEND|os.O_RDWR, 0644)

	if err != nil {
		log.Println("creating file")
		file, err = os.Create(*fileName)

		if err != nil {
			log.Fatalf("cannot read or create file %s", fileName)
		}
	}

	return file
}

func init() {
	rootCmd.AddCommand(syncCmd)

	syncCmd.Flags().String("ticker", "SPY", "Ticker to retrieve price for")
	syncCmd.Flags().String("date", "", "Date to pull historical price for")
	syncCmd.Flags().String("file", "", "file to write price to")
}
