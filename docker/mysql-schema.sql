CREATE TABLE IF NOT EXISTS `products`
(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `external_id` INT,
    `type` VARCHAR (255),
    `name` VARCHAR (255)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `products_translated`
(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `external_id` INT,
    `type` VARCHAR (255),
    `name` VARCHAR (255)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;