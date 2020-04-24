package prices

import (
	"database/sql"
	_ "database/sql"
	"encoding/csv"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
	Writer io.Writer
	DisplayToUser io.Writer
}

func (this *WriteCSV) Write(prices *TimeSeriesPrice) {
	this.writeCSV(prices)
}

func (this *WriteCSV) writeCSV(prices *TimeSeriesPrice) {
	w := csv.NewWriter(this.Writer)

	err := w.Write([]string{prices.Date, prices.Open, prices.High, prices.Low, prices.Close, prices.Volume})

	if (err != nil ) {
		log.Fatalf("Could not write to file, %s", err)
	}

	w.Flush()

	fmt.Fprintln(this.DisplayToUser, "price written")
}

type WritePostgres struct {

}

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

func (this *WritePostgres) Write(prices *TimeSeriesPrice) {
	this.writePostgres(prices)
}


func (this *WritePostgres) writePostgres(prices *TimeSeriesPrice) {
	db, err := sqlx.Connect("postgres", viper.GetString("POSTGRES_DSN"))
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()

	tx.MustExec("INSERT INTO daily_prices (date, open, high, low, close, volume) VALUEs ($1, $2, $3, $4, $5, $6)",
		prices.Date, prices.Open, prices.High, prices.Low, prices.Close, prices.Volume)

	tx.Commit()
}
