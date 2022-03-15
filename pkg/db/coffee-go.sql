CREATE TABLE coffee (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  variety VARCHAR(100),
  bitterness FLOAT,
  description VARCHAR
);

CREATE TABLE drink (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(50),
  description VARCHAR,
  photo BLOB,
  FOREIGN KEY(coffee_id) REFERENCES coffee(id)
);