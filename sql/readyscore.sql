CREATE TABLE readyscore (
    id BIGINT GENERATED ALWAYS AS IDENTITY ,
    date DATE NOT NULL,
    score INTEGER NOT NULL,
    created_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);
CREATE UNIQUE INDEX idx_readyscore_date ON readyscore(date);
ALTER TABLE readyscore ADD CONSTRAINT unique_readyscore_date UNIQUE(date);

CREATE OR REPLACE FUNCTION update_readyscore_updated_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_timestamp = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER readyscore_updated_trigger
    BEFORE UPDATE ON readyscore
    FOR EACH ROW EXECUTE FUNCTION update_readyscore_updated_timestamp();