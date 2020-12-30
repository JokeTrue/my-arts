CREATE TABLE IF NOT EXISTS categories
(
    id         INT UNSIGNED AUTO_INCREMENT NOT NULL,

    title      VARCHAR(255)                    NOT NULL,
    created_at DATETIME                        NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id),
    UNIQUE KEY title (title)
);


CREATE TABLE IF NOT EXISTS products
(
    id                  INT UNSIGNED AUTO_INCREMENT                           NOT NULL,
    user_id             INT UNSIGNED                                          NOT NULL,
    category_id         INT UNSIGNED                                          NOT NULL,


    title               VARCHAR(255)                                              NOT NULL,
    general_description TEXT,
    product_description TEXT,
    state               ENUM ('NEW', 'IN_PROGRESS', 'SOLD', 'CLOSED', 'ARCHIVED') NOT NULL,
    created_at          DATETIME                                                  NOT NULL DEFAULT NOW(),

    price_amount        DECIMAL(13, 2)                                            NOT NULL,
    price_currency      VARCHAR(3)                                                NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (category_id) REFERENCES categories (id)
);

CREATE TABLE IF NOT EXISTS product_tags
(
    id         INT UNSIGNED AUTO_INCREMENT NOT NULL,
    product_id INT UNSIGNED                NOT NULL,
    title      VARCHAR(255)                    NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (product_id) REFERENCES products (id)
);

CREATE TABLE IF NOT EXISTS product_photos
(
    id         INT UNSIGNED AUTO_INCREMENT NOT NULL,
    product_id INT UNSIGNED                NOT NULL,

    url        TEXT                            NOT NULL,
    created_at DATETIME                        NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id),
    FOREIGN KEY (product_id) REFERENCES products (id)
);



