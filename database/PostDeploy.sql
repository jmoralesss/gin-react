CREATE DATABASE IF NOT EXISTS Array;

CREATE USER 'array_user'@'localhost' IDENTIFIED BY 'password123';

GRANT ALL PRIVILEGES ON Array.* TO 'array_user'@'localhost';

USE Array

CREATE TABLE IF NOT EXISTS User (
   ID INT NOT NULL AUTO_INCREMENT,
   Email VARCHAR(100) NOT NULL,
   Password VARCHAR(256) NOT NULL,
   InsertDate DateTime NOT NULL,

   PRIMARY KEY ( ID )
);