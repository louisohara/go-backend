DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id         INT AUTO_INCREMENT NOT NULL,
  first_name      VARCHAR(50) NOT NULL,
  surname     VARCHAR(50) NOT NULL,
  email      VARCHAR(50) NOT NULL,
  date_of_birth VARCHAR(10) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO users
  (first_name, surname, email, date_of_birth)
VALUES
  ('Louis', 'OHara', 'louisohara20@gmail.com', '19/06/1999'),
  ('Simon', 'Thomas', 'simon@element3.tech', "unknown"),
  ('Carl', 'Barker', 'carl@element3.tech', 'unknown')