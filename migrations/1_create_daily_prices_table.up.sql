CREATE TABLE daily_prices (
    symbol VARCHAR(15) NOT NULL,
    date DATE NOT NULL,
    open_price NUMERIC(20, 6) NOT NULL,
    high_price NUMERIC(20, 6) NOT NULL,
    low_price NUMERIC(20, 6) NOT NULL,
    close_price NUMERIC(20, 6) NOT NULL,
    volume integer,
    CONSTRAINT symbol_date UNIQUE(symbol, date)
)
