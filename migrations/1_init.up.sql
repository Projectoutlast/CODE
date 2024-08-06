CREATE TABLE IF NOT EXISTS menu (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    menu_type VARCHAR NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS category_dish (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    menu_type_id INTEGER NOT NULL,
    category_name VARCHAR UNIQUE NOT NULL,
    FOREIGN KEY (menu_type_id) REFERENCES menu(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dishes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    dish_name VARCHAR NOT NULL,
    menu_type_id INTEGER NOT NULL,
    category_dish_id INTEGER NOT NULL,
    composition_of_the_dish VARCHAR NOT NULL,
    dish_description VARCHAR,
    price NUMERIC(4, 2) NOT NULL,
    dish_weight INTEGER NOT NULL,
    dish_image BLOB,
    tags VARCHAR ARRAY,
    FOREIGN KEY (category_dish_id) REFERENCES category_dish(id) ON DELETE CASCADE,
    FOREIGN KEY (menu_type_id) REFERENCES menu(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS admin_panel_users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_login VARCHAR NOT NULL UNIQUE,
    user_password VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    firstname VARCHAR NOT NULL,
    lastname VARCHAR,
    user_role VARCHAR NOT NULL DEFAULT 'менеджер',
    create_date TIMESTAMP DEFAULT current_timestamp,
    update_date TIMESTAMP DEFAULT current_timestamp,
    CHECK (user_role IN ('менеджер', 'управляющий'))
);