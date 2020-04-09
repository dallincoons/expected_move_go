package prices

import (
	"fmt"
	"os"
	"encoding/csv"
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
	fmt.Fprintf(os.Stdout,"%-10s | %-8s | %-8s | %-8s | %-8s | %-8s\n", "Date", "Open", "High", "Low", "Close", "Volume")
	fmt.Fprintf(os.Stdout, "%-10s | %-8s | %-8s | %-8s | %-8s | %-8s\n", prices.Date, prices.Open, prices.High, prices.Low, prices.Close, prices.Volume)
}

type WriteCSV struct {

}

func (this *WriteCSV) Write(prices *TimeSeriesPrice) {
	this.writeCSV(prices)
}

func (this *WriteCSV) writeCSV(prices *TimeSeriesPrice) {
	w := csv.NewWriter(os.Stdout)

	w.Write([]string{prices.Date, prices.Open, prices.High, prices.Low, prices.Close, prices.Volume})

	w.Flush()
}

