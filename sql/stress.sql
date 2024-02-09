CREATE TABLE stress (
    id uuid DEFAULT gen_random_uuid(),
    date DATE NOT NULL,
    high_stress_duration INTEGER NOT NULL, -- milliseconds
    created_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
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
