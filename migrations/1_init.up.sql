CREATE TABLE IF NOT EXISTS menu (
    id SERIAL PRIMARY KEY,
    menu_type VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS category (
    id SERIAL PRIMARY KEY,
    category_type VARCHAR NOT NULL,
    category_dish VARCHAR NOT NULL,
    FOREIGN KEY (category_type) REFERENCES menu(id)
);

CREATE TABLE IF NOT EXISTS dishes (
    id SERIAL PRIMARY KEY,
    dish_name VARCHAR UNIQUE NOT NULL,
    category_dish INTEGER NOT NULL,
    composition_of_the_dish VARCHAR NOT NULL,
    description VARCHAR,
    price NUMERIC(4, 2) NOT NULL,
    weight INTEGER NOT NULL,
    image BLOB,
    tags VARCHAR ARRAY,
    FOREIGN KEY (category_dish) REFERENCES category(id)
);

CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    event_name VARCHAR NOT NULL,
    description VARCHAR,
    event_date TIMESTAMP DEFAULT current_timestamp,
    event_time TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS admin_panel_users (
    id SERIAL PRIMARY KEY,
    login VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    name VARCHAR NOT NULL,
    lastname VARCHAR,
    role VARCHAR NOT NULL DEFAULT 'менеджер',
    create_date TIMESTAMP DEFAULT current_timestamp,
    update_date TIMESTAMP DEFAULT current_timestamp,
    CHECK (role IN ('менеджер', 'управляющий'))
);