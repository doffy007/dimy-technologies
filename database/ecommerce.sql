CREATE TABLE `customer` (
  `id` int NOT NULL AUTO_INCREMENT,
  `customer_name` char(50) NOT NULL,
  PRIMARY KEY(`id`)
);

CREATE TABLE `customer_address` (
  `id` int NOT NULL AUTO_INCREMENT,
  `customer_id` int NOT NULL,
  `address` varchar(255) NOT NULL,
  PRIMARY KEY(`id`),
  FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`)
);

CREATE TABLE `product` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `price` decimal NOT NULL,
  PRIMARY KEY(`id`)
);

CREATE TABLE `payment_method` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `is_active` tinyint NOT NULL,
  PRIMARY KEY(`id`)
);

CREATE TABLE `transaction` (
  `id` int NOT NULL AUTO_INCREMENT,
  `customer_address_id` int NOT NULL,
  `product_id` int NOT NULL,
  `payment_method_id` int NOT NULL,
  `product_qty` int NOT NULL,
  PRIMARY KEY(`id`),
  FOREIGN KEY (`customer_address_id`) REFERENCES `customer_address` (`id`),
  FOREIGN KEY (`product_id`) REFERENCES `product` (`id`),
  FOREIGN KEY (`payment_method_id`) REFERENCES `payment_method` (`id`)
);
