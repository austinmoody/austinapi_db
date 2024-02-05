CREATE TABLE stress (
      id SERIAL PRIMARY KEY,
      date DATE,
      high_stress_duration INTEGER, -- milliseconds
      created_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_timestamp TIMESTAMP
);

CREATE INDEX idx_stress_date ON stress(date);
ALTER TABLE stress ADD CONSTRAINT unique_stress_date UNIQUE(date);

CREATE OR REPLACE FUNCTION update_stress_updated_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_timestamp = CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER preparedness_stress_trigger
    BEFORE UPDATE ON stress
    FOR EACH ROW EXECUTE FUNCTION update_stress_updated_timestamp();
