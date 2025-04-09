-- +goose Up
CREATE TABLE measurements (
    id UUID PRIMARY KEY,
    measurement_time TIMESTAMP NOT NULL,
    humidity FLOAT NOT NULL,
    temperature FLOAT NOT NULL
);

-- +goose Down
DROP TABLE measurements;
