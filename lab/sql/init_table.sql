CREATE TABLE favorite_movies
(
    movie_id   VARCHAR PRIMARY KEY,
    title      VARCHAR,
    year       INTEGER,
    image      VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);