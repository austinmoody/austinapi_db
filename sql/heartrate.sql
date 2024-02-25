CREATE TABLE heartrate (
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    date DATE NOT NULL,
    high INTEGER NOT NULL,
    low INTEGER NOT NULL,
    average INTEGER NOT NULL,
    created_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX idx_heartrate_date ON heartrate(date);
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
