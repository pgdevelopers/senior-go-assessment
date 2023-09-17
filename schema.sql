CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(255),
  last_name VARCHAR(255),
  email VARCHAR(255)
);

INSERT INTO users (first_name, last_name, email) 
VALUES ('Josh', 'Rose', 'jrose@example.com');

INSERT INTO users (first_name, last_name, email) 
VALUES ('Erica', 'Handelman', 'ehandelman@example.com');

INSERT INTO users (first_name, last_name, email) 
VALUES ('Eric', 'Boggs', 'eboggs@example.com');

INSERT INTO users (first_name, last_name, email) 
VALUES ('Megan', 'Janes', 'mjanes@example.com');
