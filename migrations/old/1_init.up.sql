CREATE TABLE IF NOT EXISTS menu (
    menu_type VARCHAR PRIMARY KEY,
    CHECK (menu_type IN ('Основное', 'Бар', 'Ланчи', 'Кейтеринг', 'Доставка', 'Кальян'))
);

CREATE TABLE IF NOT EXISTS category (
    id SERIAL PRIMARY KEY,
    menu_type VARCHAR NOT NULL,
    category_dish VARCHAR NOT NULL,
    FOREIGN KEY (menu_type) REFERENCES menu(menu_type),
    CHECK (category_dish IN (
        'Салаты', 'Закуски', 'Супы','Основные блюда', 'Гриль', 'Завтраки', 'Чай',
           'Кофе', 'Легкий алкоголь', 'Крепкий алкоголь', 'Коктейли', 'Безалкогольные напитки')
    ),
    CHECK (category_dish IN (
        'Салаты', 'Закуски', 'Супы','Основные блюда', 'Гриль', 'Завтраки', 'Чай', 'Кофе'
          ) AND menu_type IN ('Основное', 'Ланчи', 'Кейтеринг', 'Доставка')
    ),
    CHECK (category_dish IN (
           'Легкий алкоголь', 'Крепкий алкоголь', 'Коктейли', 'Безалкогольные напитки'
          ) AND menu_type IN ('Бар')
    )
);

CREATE TABLE IF NOT EXISTS dishes (
    dish_name PRIMARY KEY,
    category_dish VARCHAR NOT NULL,
    composition_of_the_dish VARCHAR NOT NULL,
    description VARCHAR,
    price NUMERIC(4, 2) NOT NULL,
    weight INTEGER NOT NULL,
    image BLOB,
    FOREIGN KEY (category_dish) REFERENCES category(category_dish)
);