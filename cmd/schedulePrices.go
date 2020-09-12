package cmd

import (
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"time"
)

// schedulePricesCmd represents the schedulePrices command
var schedulePricesCmd = &cobra.Command{
	Use:   "schedulePrices",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		schedulePricePulling()
	},
}

func schedulePricePulling() {
	c := cron.New()

	//c.AddFunc("0 16 * * 1-5", func () {
	//	pullPrices()
	//})

	c.AddFunc("* * * * *", func () {
		pullPrices()
	})

	c.Start()

	KeepAlive()
}

func pullPrices() {
	PullPrices(time.Now().AddDate(0,0,0).Format("2006-01-02"), time.Now().Format("2006-01-02"))
}

func init() {
	rootCmd.AddCommand(schedulePricesCmd)
}
