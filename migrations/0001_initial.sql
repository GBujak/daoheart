-- +goose Up
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    username VARCHAR(255) NOT NULL
);