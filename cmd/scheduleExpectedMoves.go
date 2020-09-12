package cmd

import (
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
)

func NewScheduleExpectedMovesCmd() *cobra.Command {
 return &cobra.Command{
		Use:   "scheduleExpectedMoves",
		Short: "",
		Long: ``,
		Run: func(cmd *cobra.Command, args []string) {
			ScheduleExpectedMovePulling()
		},
	}
}

func ScheduleExpectedMovePulling() {
	c := cron.New()

	c.AddFunc("0 0 * * 6", func() {
		pullExpectedMoves();
	})

	c.Start()

	KeepAlive()
}

func pullExpectedMoves() {
	SyncExpectedMoves()
}

func init() {
	rootCmd.AddCommand(NewScheduleExpectedMovesCmd())
}
