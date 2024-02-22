CREATE TABLE preparedness (
    id BIGINT GENERATED ALWAYS AS IDENTITY ,
    date DATE NOT NULL,
    rating INTEGER NOT NULL,
    created_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);
CREATE INDEX idx_preparedness_date ON preparedness(date);
ALTER TABLE preparedness ADD CONSTRAINT unique_preparedness_date UNIQUE(date);

CREATE OR REPLACE FUNCTION update_preparedness_updated_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_timestamp = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER preparedness_updated_trigger
    BEFORE UPDATE ON preparedness
    FOR EACH ROW EXECUTE FUNCTION update_preparedness_updated_timestamp();