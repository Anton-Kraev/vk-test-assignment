-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS container (
    id SERIAL PRIMARY KEY,
    ip INET NOT NULL UNIQUE,
    last_ping_attempt TIMESTAMP WITH TIME ZONE,
    last_successful_ping TIMESTAMP WITH TIME ZONE,
    response_time_ms INT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS container;
-- +goose StatementEnd
