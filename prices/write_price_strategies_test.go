package prices

import (
	"bytes"
	"testing"
)

func TestWriteToCSV(t *testing.T) {
	writer := &FakeWriter{
		Output: &bytes.Buffer{},
	}
	fakeStdOut := &FakeWriter{
		Output: &bytes.Buffer{},
	}

	csvWriter := &WriteCSV{
		Writer: writer,
		DisplayToUser: fakeStdOut,
	}

	csvWriter.writeCSV(&TimeSeriesPrice{
		Date:   "2020-01-01",
		Open:   "101.13",
		High:   "101.14",
		Low:    "101.15",
		Close:  "101.16",
		Volume: "1010101",
	})

	if (writer.Output.String() != "2020-01-01,101.13,101.14,101.15,101.16,1010101\n") {
		t.Errorf("Error writing csv contents, recieved %s, expected %s", writer.Output.String(), "020-01-01,101.13,101.14,101.15,101.16,1010101")
	}
}

var schema = `
CREATE TABLE IF NOT EXISTS daily_prices (
    date date,
    open text,
    high text,
    low text,
    close text,
    volume text 
)`

type DbTimeSeriesPrice struct {
	Date string `db:"date"`
	Open string `db:"open"`
	High string `db:"high"`
	Low string `db:"low"`
	Close string `db:"close"`
	Volume string `db:"volume"`
}

func TestWriteToPostgres(t *testing.T) {
	//db, err := sqlx.Connect("postgres", "user=postgres password=secret dbname=postgres sslmode=disable")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//db.MustExec(schema)
	//db.MustExec("DELETE from daily_prices")
	//
	//writer := &WritePostgres{}
	//
	//writer.writePostgres(&TimeSeriesPrice{
	//	Date:   "2020-01-01",
	//	Open:   "101.13",
	//	High:   "101.14",
	//	Low:    "101.15",
	//	Close:  "101.16",
	//	Volume: "1010101",
	//})
	//
	//result := []DbTimeSeriesPrice{}
	//db.Select(&result, "SELECT * FROM daily_prices")
	//
	//if result[0].Date != "2020-01-01T00:00:00Z" {
	//	t.Errorf("Date not written to database")
	//}
	//
	//if result[0].Open != "101.13" {
	//	t.Errorf("Open price not written to database")
	//}
	//
	//if result[0].High != "101.14" {
	//	t.Errorf("High price not written to database")
	//}
	//
	//if result[0].Low != "101.15" {
	//	t.Errorf("Low price not written to database")
	//}
	//
	//if result[0].Close != "101.16" {
	//	t.Errorf("Close price not written to database")
	//}
	//
	//if result[0].Volume != "1010101" {
	//	t.Errorf("Volume not written to database")
	//}
}

type FakeWriter struct {
	Output *bytes.Buffer
}

func (this *FakeWriter) Write(p []byte) (n int, err error) {
	this.Output.Write(p)

	return 0, nil
}
