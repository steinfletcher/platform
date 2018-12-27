-- +goose Up
CREATE TABLE users
(
  id       SERIAL PRIMARY KEY,
  username TEXT, -- todo create index
  password TEXT
);

INSERT INTO users (username, password)
VALUES ('zeus@me.com', '$2a$08$m.aBORIHH2.Ks67rNKWcmus2ftTqFbbwQWNYrudSHvCsXOBtvM4jW');

-- +goose Down
DROP TABLE users;
