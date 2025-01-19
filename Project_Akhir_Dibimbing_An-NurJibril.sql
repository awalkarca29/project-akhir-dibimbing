create database project_akhir_dibimbing;
drop database project_akhir_dibimbing;
use project_akhir_dibimbing;

INSERT INTO roles (name, created_at, updated_at) 
VALUES ("Admin", NOW(), NOW()), ("Customer", NOW(), NOW());

INSERT INTO products (name, description, slug, price, stock, created_at, updated_at) 
VALUES 
	("Sendal", "Sendal harian", "sendal", "10000", "5", NOW(), NOW()), 
	("Sepatu", "Sepatu harian", "sepatu", "50000", "10", NOW(), NOW()),
	("Jam", "Jam harian", "jam", "15000", "8", NOW(), NOW());

CREATE TABLE `roles` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `users` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `role_id` integer,
  `name` varchar(255),
  `email` varchar(255),
  `password` varchar(255),
  `photo` varchar(255),
  `phone` varchar(255),
  `address` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `products` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(255),
  `description` varchar(255),
  `slug` varchar(255),
  `price` integer,
  `stock` integer,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `product_images` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `product_id` integer,
  `file_name` varchar(255),
  `is_primary` integer,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `transactions` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` integer,
  `product_id` integer,
  `quantity` integer,
  `total` integer,
  `payment_method` varchar(255),
  `status` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

ALTER TABLE `users` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

ALTER TABLE `product_images` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

ALTER TABLE `transactions` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

ALTER TABLE `transactions` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
