CREATE TABLE IF NOT EXISTS admin_panel_users (
    id SERIAL PRIMARY KEY,
    login VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    name VARCHAR NOT NULL,
    lastname VARCHAR,
    role VARCHAR NOT NULL DEFAULT 'manager',
    create_date TIMESTAMP DEFAULT current_timestamp,
    update_date TIMESTAMP,
    CHECK ( role IN ('boss', 'manager') )
);