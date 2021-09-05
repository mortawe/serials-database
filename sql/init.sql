CREATE TABLE person
(
    person_id SERIAL PRIMARY KEY,
    name      VARCHAR(100),
    birthdate DATE NOT NULL,
    bio       TEXT
);

CREATE TABLE genre
(
    genre_id SERIAL PRIMARY KEY,
    name     VARCHAR(100) NOT NULL
);

CREATE TABLE show
(
    show_id     SERIAL PRIMARY KEY,
    title       VARCHAR(100),
    release     DATE NOT NULL,
    description TEXT NOT NULL,
    episode_num INT DEFAULT 0,
    genre_id    INT
        CONSTRAINT genre_id REFERENCES genre
);



CREATE TABLE person_show
(
    person_id INT
        CONSTRAINT person_id
            REFERENCES person,
    show_id   INT
        CONSTRAINT show_id
            REFERENCES show,
    CONSTRAINT person_show_uq UNIQUE (person_id, show_id)
);

