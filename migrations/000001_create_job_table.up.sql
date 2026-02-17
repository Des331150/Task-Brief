CREATE TABLE IF NOT EXISTS jobs(
    id BIGSERIAL PRIMARY KEY,
    audio_id BIGINT,
    current_status TEXT CHECK(current_status IN('queued', 'done')),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_jobs_modtime 
BEFORE UPDATE ON jobs 
FOR EACH ROW EXECUTE PROCEDURE update_modified_column();