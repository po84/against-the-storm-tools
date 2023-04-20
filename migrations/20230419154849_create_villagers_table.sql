-- +goose Up
CREATE TABLE IF NOT EXISTS villagers (
  identifier TEXT PRIMARY KEY,
  display_name TEXT
);

INSERT INTO villagers (identifier, display_name)
VALUES
  ('human', 'Human'),
  ('beaver', 'Beaver'),
  ('lizard', 'Lizard'),
  ('harpy', 'Harpy'),
  ('fox', 'Fox');

-- +goose Down
DROP TABLE IF EXISTS villagers;
