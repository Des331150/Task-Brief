/* remember to check the current status field for constraints and solve the updated at field to properly receive updates */
CREATE TABLE IF NOT EXISTS job(
    id BIGSERIAL PRIMARY KEY,
    audio_id BIGINT,
    current_status CHECK,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);