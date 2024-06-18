CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    event_name VARCHAR NOT NULL,
    description VARCHAR,
    event_date DATE NOT NULL,
    event_time TIME NOT NULL
);