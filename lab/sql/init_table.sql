CREATE TABLE favorite_movies
(
    movie_id   VARCHAR PRIMARY KEY,
    title      VARCHAR,
    year       INTEGER,
    rating     FLOAT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);