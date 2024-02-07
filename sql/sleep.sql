create table sleep
(
    id SERIAL PRIMARY KEY,
    date DATE NOT NULL,
    rating BIGINT NOT NULL,
    total_sleep INTEGER NOT NULL,
    deep_sleep INTEGER NOT NULL,
    light_sleep INTEGER NOT NULL,
    rem_sleep INTEGER NOT NULL,
    created_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_sleep_date ON sleep(date);
ALTER TABLE sleep ADD CONSTRAINT unique_sleep_date UNIQUE(date);

CREATE OR REPLACE FUNCTION update_sleep_updated_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_timestamp = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER sleep_updated_trigger
BEFORE UPDATE ON sleep
FOR EACH ROW EXECUTE FUNCTION update_sleep_updated_timestamp();