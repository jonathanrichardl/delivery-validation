CREATE DATABASE order-validator;

CREATE TABLE orders(
id integer PRIMARY KEY,
title varchar(50)
);

CREATE TABLE requirements(
requirementid integer PRIMARY KEY AUTO_INCREMENT,
request varchar(50),
expectedoutcome varchar(50),
order_id integer,
status bool,
FOREIGN KEY(order_id) REFERENCES orders(id)
);

