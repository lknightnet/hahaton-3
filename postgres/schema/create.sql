CREATE TABLE users
(
    id        SERIAL PRIMARY KEY,
    name      VARCHAR(500) NOT NULL,
    email     VARCHAR(500) NOT NULL UNIQUE,
    password  VARCHAR(500) NULL,
    status    BOOLEAN NULL,
    uuid      VARCHAR(500) NOT NULL,
    type_user VARCHAR(500) NOT NULL,
    createdAt time         NOT NULL,
    updatedAt time         NOT NULL
);

CREATE TABLE tokens
(
    id      serial primary key,
    user_id int           NOT NULL,
    token   VARCHAR(1500) NOT NULL UNIQUE
);

CREATE TABLE resumes
(
    id            SERIAL PRIMARY KEY,
    user_id       INT          NOT NULL,
    name          VARCHAR(255) NOT NULL,
    user_name     VARCHAR(255) NOT NULL,
    user_age      INT          NOT NULL,
    city          VARCHAR(255),
    university    VARCHAR(255),
    course        VARCHAR(255),
    course_number INT,
    status        BOOLEAN DEFAULT TRUE,
    description   TEXT,
    document_id   INT
);

CREATE TABLE students
(
    id         SERIAL PRIMARY KEY,
    user_id    INT          NOT NULL,
    avatar_url VARCHAR(255),
    email      VARCHAR(255) NOT NULL
);

