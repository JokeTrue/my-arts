CREATE TABLE IF NOT EXISTS friendship_requests
(
    id         INT UNSIGNED AUTO_INCREMENT NOT NULL,
    actor_id   INT UNSIGNED                NOT NULL REFERENCES users (id),
    friend_id  INT UNSIGNED                NOT NULL REFERENCES users (id),
    created_at DATETIME                    NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id),
    UNIQUE KEY friendship_uniq (actor_id, friend_id)
);