CREATE TABLE IF NOT EXISTS cryptocompare_pairs(
    id BIGSERIAL PRIMARY KEY,
    date_added TIMESTAMP DEFAULT NOW(),
    raw JSONB
);

