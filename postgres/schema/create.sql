CREATE TABLE users
(
    name      VARCHAR(500) NOT NULL,
    email     VARCHAR(500) NOT NULL UNIQUE,
    password  VARCHAR(500) NULL,
    status    BOOLEAN NULL,
    uuid      VARCHAR(500) NOT NULL,
    type_user VARCHAR(500)  NOT NULL,
    createdAt time         NOT NULL,
    updatedAt time         NOT NULL

);

CREATE TABLE tokens
(
    user_id int          NOT NULL,
    token   VARCHAR(1500) NOT NULL UNIQUE
);
