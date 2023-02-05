DROP DATABASE IF EXISTS assetmanager;
CREATE DATABASE assetmanager;
USE assetmanager;

CREATE TABLE IF NOT EXISTS `asset` (
    `id` char(36) NOT NULL UNIQUE,
    `name` varchar(255) NOT NULL,
    `amount` int NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;