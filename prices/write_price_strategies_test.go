package prices

import (
	"testing"
)

func TestWriteToCSV(t *testing.T) {
	writer := &FakeWriter{}

	csvWriter := &WriteCSV{
		Writer: writer,
	}

	csvWriter.writeCSV(&TimeSeriesPrice{
		Date:   "2020-01-01",
		Open:   "101.13",
		High:   "101.14",
		Low:    "101.15",
		Close:  "101.16",
		Volume: "1010101",
	})

	if (string(writer.Output) != "2020-01-01,101.13,101.14,101.15,101.16,1010101\n") {
		t.Errorf("Error writing csv contents, recieved %s, expected %s", writer.Output, "020-01-01,101.13,101.14,101.15,101.16,1010101")
	}
}

type FakeWriter struct {
	Output []byte
}

func (this *FakeWriter) Write(p []byte) (n int, err error) {
	this.Output = p

	return 0, nil
}
