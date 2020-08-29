package prices

import (
	"database/sql"
	_ "database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DisplayStrategyInterface interface {
	Write(prices *TimeSeriesPrice) (error)
}

type DisplayTable struct {

}

func (this *DisplayTable) Write(prices *TimeSeriesPrice) (error) {
	return this.displayTable(prices)
}

func (this *DisplayTable) displayTable(prices *TimeSeriesPrice) (error) {
	fmt.Fprintf(os.Stdout,"%-10s | %-8s | %-8s | %-8s | %-8s | %-8s\n", "Date", "Open", "High", "Low", "Close", "Volume")
	fmt.Fprintf(os.Stdout, "%-10s | %-8s | %-8s | %-8s | %-8s | %-8s\n", prices.Date, prices.Open, prices.High, prices.Low, prices.Close, prices.Volume)

	return nil
}

type WriteCSV struct {
	Writer io.ReadWriter
	DisplayToUser io.Writer
}

func (this *WriteCSV) Write(prices *TimeSeriesPrice) (error) {
	return this.writeCSV(prices)
}

func (this *WriteCSV) writeCSV(prices *TimeSeriesPrice) (error) {
	if this.recordAlreadyRecord(prices.Date) == true {
		return errors.New(fmt.Sprintf("Historical price for %s has already been written", prices.Date))
	}

	w := csv.NewWriter(this.Writer)

	writeErr := w.Write([]string{prices.Date, prices.Open, prices.High, prices.Low, prices.Close, prices.Volume})

	if (writeErr != nil ) {
		log.Fatalf("Could not write to file, %s", writeErr)
	}

	w.Flush()

	werr := w.Error()

	if (werr != nil) {
		log.Fatal(werr)
	}

	fmt.Fprintln(this.DisplayToUser, "price written")

	return nil
}

// For now, use a simple brute force approach
func (this *WriteCSV) recordAlreadyRecord(date string) bool {
	r := csv.NewReader(this.Writer)

	var record []string

	for {
		r, err := r.Read()
		record = r

		if (err == io.EOF) {
			break
		}
		if (err != nil) {
			log.Fatal(err)
		}

		if (record[0] == date) {
			return true
		}
	}

	return false
}

type WritePostgres struct {
	Dsn string
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

func (this *WritePostgres) Write(prices *TimeSeriesPrice) (error) {
	if prices == nil {
		log.Fatal("Price not found")
	}

	db, err := sqlx.Connect("postgres", this.Dsn)
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()

	query := `
		INSERT INTO daily_prices
		(symbol, date, open_price, high_price, low_price, close_price, volume)	
		 VALUES
		 ($1, $2, $3, $4, $5, $6, $7)
		 ON CONFLICT ON CONSTRAINT symbol_date
		 DO
		 	UPDATE SET open_price = EXCLUDED.open_price, high_price = EXCLUDED.high_price, 
		 				low_price = EXCLUDED.low_price, close_price = EXCLUDED.close_price, 
		 				volume = EXCLUDED.volume
	`

	//todo: handle errors
	tx.MustExec(query, prices.Ticker, prices.Date, prices.Open, prices.High, prices.Low, prices.Close, prices.Volume)

	err = tx.Commit()

	if err != nil {
		log.Print(err)
	}

	return nil
}
