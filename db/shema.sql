CREATE TABLE user (
    id       INTEGER PRIMARY KEY AUTOINCREMENT
                     NOT NULL,
    login    TEXT    NOT NULL,
    password TEXT    NOT NULL,
    token    TEXT,
    option   BLOB
);
