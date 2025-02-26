CREATE DATABASE IF NOT EXISTS `devbook`;
Use `devbook`;

Drop table if exists `users`;
CREATE TABLE `users` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(50) NOT NULL,
    `email` VARCHAR(50) NOT NULL,
    `password` VARCHAR(50) NOT NULL
    `nickname` VARCHAR(50) NOT NULL unique,
    `crated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)Engine=InnoDB;

CREATE TABLE if not exists `userAdress` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `cep` VARCHAR(9) NOT NULL,
    `street` VARCHAR(50) NOT NULL,
    `number` INT NOT NULL,
    `complement` VARCHAR(50),
    `city` VARCHAR(50) NOT NULL,
    `state` VARCHAR(50) NOT NULL,
    `contry` VARCHAR(50) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
)Engine=InnoDB;

CREATE TABLE if not exists `userConfig` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `theme` VARCHAR(50) NOT NULL,
    `language` VARCHAR(50) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
)Engine=InnoDB;