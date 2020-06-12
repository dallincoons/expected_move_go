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
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

// syncAllCmd represents the syncAll command
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
		filepath, _ := cmd.Flags().GetString("filepath")

		pricesController := &prices.HistoricalPriceController{
			Client: http.DefaultClient,
		}

		writers := createWriters(filepath);

		pricesController.GetPrices(date, writers)
	},
}

func createWriters(filepath string) map[string]prices.DisplayStrategyInterface {
	writers := make(map[string]prices.DisplayStrategyInterface)
	for _, ticker := range prices.GetTickers() {
		writers[ticker] = prices.GetWriteStrategy(fmt.Sprintf("%s/%s.csv", filepath, ticker))
	}

	return writers
}

func init() {
	rootCmd.AddCommand(syncAllCmd)

	syncAllCmd.Flags().String("date", time.Now().Format("2006-01-02"), "Date to pull historical prices for")
}
