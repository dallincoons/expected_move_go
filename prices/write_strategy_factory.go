package prices

import (
	"log"
	"os"
)

func GetWriteStrategy(fileName string) DisplayStrategyInterface {
	if (fileName != "") {
		return &WriteCSV{
			Writer:        getCsvWriter(&fileName),
			DisplayToUser: os.Stdout,
		}
	}

	return &DisplayTable{}
}

func getCsvWriter(filePath *string) *os.File {
	file, err := os.OpenFile(*filePath, os.O_APPEND|os.O_RDWR, 0644)

	if err != nil {
		log.Println("creating file")
		file, err = os.Create(*filePath)

		if err != nil {
			log.Fatalf("cannot read or create file %s", *filePath)
		}
	}

	return file
}
