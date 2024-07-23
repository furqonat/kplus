-- Active: 1721762008159@@127.0.0.1@3306@kplus
CREATE TABLE IF NOT EXISTS `users` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `phone` varchar(255) NOT NULL UNIQUE,
    `email` varchar(255) UNIQUE,
    `role` ENUM('admin', 'user') NOT NULL DEFAULT 'user',
    `phone_verified` BOOLEAN NOT NULL DEFAULT FALSE,
    `email_verified` BOOLEAN NOT NULL DEFAULT FALSE,
    `password` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `user_details` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_id` int(11) NOT NULL,
    `identity_number` varchar(255) NOT NULL UNIQUE,
    `full_name` varchar(255) NOT NULL,
    `legal_name` varchar(255) NOT NULL,
    `place_of_birth` varchar(255) NOT NULL,
    `date_of_birth` date NOT NULL,
    `salary` DECIMAL NOT NULL,
    `selfie` varchar(255) NOT NULL,
    `selfie_with_national_id` varchar(255) NOT NULL,
    `national_id_image` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `loan_limits` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_id` int(11) NOT NULL,
    `limit` DECIMAL NOT NULL,
    `tenor` int(4) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `transactions` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `contract_number` varchar(255) NOT NULL,
    `user_id` int(11) NOT NULL,
    `otr` DECIMAL NOT NULL,
    `fee` DECIMAL NOT NULL,
    `installment` DECIMAL NOT NULL,
    `interest` DECIMAL NOT NULL,
    `status` ENUM(
        'pending',
        'success',
        'failed'
    ) NOT NULL DEFAULT 'pending',
    `asset_name` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `installments` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `transaction_id` int(11) NOT NULL,
    `installment` DECIMAL NOT NULL,
    `due_date` date NOT NULL,
    `paid_date` date,
    `period` int(4) NOT NULL,
    `status` ENUM('unpaid', 'paid', 'due') NOT NULL DEFAULT 'unpaid',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`) ON DELETE CASCADE
);