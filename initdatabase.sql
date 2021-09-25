use testers;
CREATE TABLE orders(
id integer PRIMARY KEY AUTO_INCREMENT,
title varchar(50)
);

CREATE TABLE requirements(
question varchar(50),
answer varchar(50),
order_id integer,
status bool,
FOREIGN KEY(order_id) REFERENCES orders(id)
);

