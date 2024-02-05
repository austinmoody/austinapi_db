CREATE TABLE heartrate (
    id SERIAL PRIMARY KEY,
    date DATE,
    high INTEGER,
    low INTEGER,
    average INTEGER,
    created_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_timestamp TIMESTAMP
);

CREATE INDEX idx_heartrate_date ON heartrate(date);
ALTER TABLE heartrate ADD CONSTRAINT unique_heartrate_date UNIQUE(date);

CREATE OR REPLACE FUNCTION update_heartrate_updated_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_timestamp = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER preparedness_heartrate_trigger
    BEFORE UPDATE ON heartrate
    FOR EACH ROW EXECUTE FUNCTION update_heartrate_updated_timestamp();
