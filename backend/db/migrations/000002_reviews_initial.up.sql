CREATE TABLE IF NOT EXISTS reviews
(
    id                   INT UNSIGNED AUTO_INCREMENT NOT NULL,
    user_id              INT UNSIGNED                NOT NULL,
    reviewer_id          INT UNSIGNED                NOT NULL,

    comment              TEXT,
    delivery_rating      TINYINT(1)                  NOT NULL DEFAULT 0,
    communication_rating TINYINT(1)                  NOT NULL DEFAULT 0,
    accuracy_rating      TINYINT(1)                  NOT NULL DEFAULT 0,


    edited_at            DATETIME,
    created_at           DATETIME                    NOT NULL DEFAULT NOW(),


    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (reviewer_id) REFERENCES users (id)
);