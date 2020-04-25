package prices

import (
	"bytes"
	"io"
	"testing"
)

func TestWriteToCSV(t *testing.T) {
	writer := new(bytes.Buffer)

	csvWriter := &WriteCSV{
		Writer: writer,
		DisplayToUser: new(bytes.Buffer),
	}

	csvWriter.writeCSV(&TimeSeriesPrice{
		Date:   "2020-01-01",
		Open:   "101.13",
		High:   "101.14",
		Low:    "101.15",
		Close:  "101.16",
		Volume: "1010101",
	})

	if (writer.String() != "2020-01-01,101.13,101.14,101.15,101.16,1010101\n") {
		t.Errorf("Error writing csv contents, recieved %s, expected %s", writer.String(), "020-01-01,101.13,101.14,101.15,101.16,1010101")
	}
}

func TestWriteToCsvDoesNotOverwriteTheLastRecordIfDatesMatch(t *testing.T) {
	writer := new(bytes.Buffer)

	writer.WriteString("2020-01-01,101.13,101.14,101.15,101.16,1010101\n2020-01-02,101.13,101.14,101.15,101.16,1010101")

	csvWriter := &WriteCSV{
		Writer: writer,
		DisplayToUser: new(bytes.Buffer),
	}

	err := csvWriter.writeCSV(&TimeSeriesPrice{
		Date:   "2020-01-01",
		Open:   "101.13",
		High:   "101.14",
		Low:    "101.15",
		Close:  "101.16",
		Volume: "1010101",
	})

	if (err == nil) {
		t.Fatalf("Expected error due to price already having been written.")
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
	TimesRead int
	Output *bytes.Buffer
	Contents []byte
}

func (this *FakeWriter) Write(p []byte) (n int, err error) {
	this.Output.Write(p)

	return 0, nil
}

func (this *FakeWriter) Read(p []byte) (n int, err error) {
	if this.TimesRead > 0 {
		return 0, io.EOF
	}

	p = append(p, this.Contents...)

	this.TimesRead++

	return len(this.Contents), nil
}

func (this *FakeWriter) Close() error {
	return nil
}
