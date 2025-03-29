DROP DATABASE IF EXISTS `cars_db`;

CREATE DATABASE `cars_db`;

USE `cars_db`;

CREATE TABLE IF NOT EXISTS cars (
    id VARCHAR(8) PRIMARY KEY,
    model VARCHAR(100) NOT NULL
    make VARCHAR(100) NOT NULL
    package VARCHAR(100) NOT NULL
    color VARCHAR(100) NOT NULL
    fabricationYear INT NOT NULL
    category VARCHAR(100) NOT NULL
    mileage INT NOT NULL
    price INT NOT NULL
);