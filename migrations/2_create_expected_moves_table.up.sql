CREATE TABLE expected_moves (
    symbol VARCHAR(15) NOT NULL,
    week_start DATE NOT NULL,
    week_end DATE NOT NULL,
    start_price NUMERIC(20, 6) NOT NULL,
    high_price NUMERIC(20, 6) NOT NULL,
    low_price NUMERIC(20, 6) NOT NULL,
    CONSTRAINT symbol_week_start UNIQUE(symbol, week_start)
)
