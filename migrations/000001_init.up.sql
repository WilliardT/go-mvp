CREATE SCHEMA go_mvp_app;

CREATE TABLE go_mvp_app.users (
    id            SERIAL                    PRIMARY KEY,
    version       BIGINT         NOT NULL   DEFAULT 1,
    full_name     VARCHAR(100)   NOT NULL   CHECK (char_length(full_name) BETWEEN 3 AND 100),
    phone_number  VARCHAR(15)               CHECK (
        phone_number ~ '^\+[0-9]+$'
        AND 
        char_length(phone_number) BETWEEN 10 AND 15
    )
);

CREATE TABLE go_mvp_app.products (
    id              SERIAL                     PRIMARY KEY,
    version         BIGINT          NOT NULL   DEFAULT 1,
    title           VARCHAR(100)    NOT NULL   CHECK (char_length(title) BETWEEN 1 AND 100),
    description     TEXT                       CHECK (char_length(description) BETWEEN 1 AND 100),
    price           DECIMAL(10, 2)  NOT NULL   CHECK (price > 0),
    created_at      TIMESTAMPTZ,
    updated_at      TIMESTAMPTZ,    

    author_user_id  INTEGER         NOT NULL   REFERENCES go_mvp_app.users(id)
);
