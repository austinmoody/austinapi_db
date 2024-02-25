CREATE TABLE spo2 (
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    date DATE NOT NULL,
    average_spo2 DOUBLE PRECISION NOT NULL,
    created_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX idx_spo2_date ON spo2(date);
ALTER TABLE spo2 ADD CONSTRAINT unique_spo2_date UNIQUE(date);

CREATE OR REPLACE FUNCTION update_spo2_updated_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_timestamp = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER preparedness_spo2_trigger
    BEFORE UPDATE ON spo2
    FOR EACH ROW EXECUTE FUNCTION update_spo2_updated_timestamp();
