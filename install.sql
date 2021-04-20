DROP DATABASE IF EXISTS codingOrganizer;
CREATE DATABASE codingOrganizer;
USE codingOrganizer;

CREATE TABLE entries (

  entry_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) DEFAULT "",
  url VARCHAR(2083) DEFAULT "",
  codeblock TEXT,
  note TEXT

);