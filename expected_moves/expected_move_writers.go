package expected_moves

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type EMWriter interface {
	Write(moves []ExpectedMove) error
}

type PostgresWriter struct {
	Dsn string
}

func (w PostgresWriter) Write(moves []ExpectedMove) error {

	db, err := sqlx.Connect("postgres", w.Dsn)
	defer db.Close()

	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()

	for _, move := range moves {
		query := `
		INSERT INTO expected_moves
		(symbol, week_start, week_end, start_price, high_price, low_price)	
		 VALUES
		 ($1, $2, $3, $4, $5, $6)
		 ON CONFLICT ON CONSTRAINT symbol_week_start
		 DO
		 	UPDATE SET week_end = EXCLUDED.week_end, start_price = EXCLUDED.start_price, 
		 				high_price = EXCLUDED.high_price, low_price = EXCLUDED.low_price
		`
		tx.MustExec(query, move.Symbol, move.PeriodStartDate, move.PeriodEndDate, move.StartPrice, move.HighPrice, move.LowPrice)
	}

	return tx.Commit()
}
