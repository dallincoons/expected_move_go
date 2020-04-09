package prices

import (
	"fmt"
	"os"
)

type DisplayStrategyInterface interface {
	Write(prices *TimeSeriesPrice)
}

type DisplayTable struct {

}

func (this *DisplayTable) Write(prices *TimeSeriesPrice) {
	this.displayTable(prices)
}

func (this *DisplayTable) displayTable(prices *TimeSeriesPrice) {
	open := truncatePrice(prices.Open)
	high := truncatePrice(prices.High)
	low := truncatePrice(prices.Low)
	close := truncatePrice(prices.Close)

	fmt.Fprintf(os.Stdout,"%-10s | %-8s | %-8s | %-8s | %-8s | %-8s\n", "Date", "Open", "High", "Low", "Close", "Volume")
	fmt.Fprintf(os.Stdout, "%-10s | %-8s | %-8s | %-8s | %-8s | %-8s\n", prices.Date, open, high, low, close, prices.Volume)
}
