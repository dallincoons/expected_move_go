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
	"expected_move/expected_moves"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

// syncExpectedMoveCmd represents the syncExpectedMove command
var syncExpectedMoveCmd = &cobra.Command{
	Use:   "syncExpectedMoves",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
		   and usage of using your command. For example:

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		SyncExpectedMoves()
	},
}

func SyncExpectedMoves() {
	puller := expected_moves.ExpectedMovePuller{
		HttpClient: expected_moves.HttpClient{
			Client: http.Client{},
			ApiKey: os.Getenv("TDA_API_KEY"),
		},
		WeekendFinder: expected_moves.WeekendFinder{
			Calendar: expected_moves.Calendar{},
		},
	}

	moves := puller.GetExpectedMoves()

	writer := expected_moves.PostgresWriter{
		Dsn: fmt.Sprintf("postgresql://%s:%s@postgres_db:5432/%s?sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DATABASE"),
		),
	}

	writer.Write(moves)
}

func init() {
	godotenv.Load()

	rootCmd.AddCommand(syncExpectedMoveCmd)
}
