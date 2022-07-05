-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `password` longtext NOT NULL
);
-- CREATE TABLE users (
--   id bigint(20) PRIMARY KEY,
--   name varchar(255) NOT NULL,
--   email varchar(255) UNIQUE NOT NULL,
--   password text NOT NULL,
-- );

CREATE TABLE `books` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `title` varchar(100),
  `description` varchar(255),
  `user_id` int 
);

ALTER TABLE `books` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);