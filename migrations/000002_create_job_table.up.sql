CREATE EXTENSION moddate;

CREATE TABLE IF NOT EXISTS jobs(
    id BIGSERIAL PRIMARY KEY,
    audio_id BIGINT,
    current_status TEXT CHECK(current_status IN('queued', 'done')),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

CREATE TRIGGER mdt_table
    BEFORE UPDATE ON jobs
    FOR EACH ROW
    EXECUTE PROCEDURE moddate (updated_at);
