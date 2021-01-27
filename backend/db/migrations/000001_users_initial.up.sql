CREATE TABLE IF NOT EXISTS users
(
    id          INT UNSIGNED AUTO_INCREMENT NOT NULL,
    email       VARCHAR(255)                NOT NULL,
    password    VARCHAR(255)                NOT NULL,

    first_name  VARCHAR(255)                NOT NULL,
    last_name   VARCHAR(255)                NOT NULL,
    age         INT                         NOT NULL,
    gender      VARCHAR(1)                  NOT NULL,
    location    VARCHAR(255)                NOT NULL,
    biography   TEXT,

    permissions TEXT,
    created_at  DATETIME                    NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id),
    UNIQUE KEY email (email)
);


CREATE TABLE friendships
(
    user_1 INT UNSIGNED NOT NULL REFERENCES users (id),
    user_2 INT UNSIGNED NOT NULL REFERENCES users (id),

    CONSTRAINT CheckOneWay CHECK (user_1 < user_2),
    CONSTRAINT PK_Friends_UserID1_UserID2 PRIMARY KEY (user_1, user_2),
    CONSTRAINT UQ_Friends_UserID2_UserID1 UNIQUE (user_2, user_1)
);
