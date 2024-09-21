CREATE TABLE users
(
    name      VARCHAR(255) NOT NULL,
    email     VARCHAR(255) NOT NULL UNIQUE,
    password  VARCHAR(255) NULL,
    status    BOOLEAN NULL,
    uuid      VARCHAR(255) NOT NULL,
    type_user VARCHAR(50)  NOT NULL,
    createdAt time         NOT NULL,
    updatedAt time         NOT NULL

);

CREATE TABLE tokens
(
    user_id int          NOT NULL,
    token   VARCHAR(255) NOT NULL UNIQUE
);
