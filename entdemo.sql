SELECT @@version, @@collation_server, @@character_set_server;

SELECT
    `SCHEMA_NAME`, `DEFAULT_CHARACTER_SET_NAME`, `DEFAULT_COLLATION_NAME`
from
    `INFORMATION_SCHEMA`.`SCHEMATA`
WHERE `SCHEMA_NAME` = SCHEMA()
ORDER BY `SCHEMA_NAME`;


SELECT
    t1.TABLE_SCHEMA,
    t1.TABLE_NAME,
    t2.CHARACTER_SET_NAME,
    t1.TABLE_COLLATION,
    t1.AUTO_INCREMENT,
    t1.TABLE_COMMENT,
    t1.CREATE_OPTIONS
FROM
    INFORMATION_SCHEMA.TABLES AS t1
        JOIN INFORMATION_SCHEMA.COLLATIONS AS t2
             ON t1.TABLE_COLLATION = t2.COLLATION_NAME
WHERE
        TABLE_SCHEMA IN ('entdemo')
  AND TABLE_NAME IN ('cars','groups','users','group_users')
ORDER BY
    TABLE_SCHEMA, TABLE_NAME;


CREATE TABLE `users` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT,
     `age` bigint(20) NOT NULL,
     `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
     `rank` double DEFAULT NULL,
     `active` tinyint(1) NOT NULL DEFAULT '0',
     `created_at` timestamp NOT NULL,
     `url` json DEFAULT NULL,
     `strings` json DEFAULT NULL,
     `state` enum('on','off') COLLATE utf8mb4_bin DEFAULT NULL,
     `uuid` char(36) COLLATE utf8mb4_bin NOT NULL,
     PRIMARY KEY (`id`),
     UNIQUE KEY `name` (`name`),
     UNIQUE KEY `user_age_name` (`age`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE `cars` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `model` varchar(255) NOT NULL,
    `registered_at` timestamp NOT NULL,
    `user_cars` bigint NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `cars_users_cars` FOREIGN KEY (`user_cars`) REFERENCES `users` (`id`) ON DELETE SET NULL
) CHARSET utf8mb4 COLLATE utf8mb4_bin;

CREATE TABLE `groups` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;


CREATE TABLE `group_users` (
    `group_id` bigint NOT NULL,
    `user_id` bigint NOT NULL,
    PRIMARY KEY (`group_id`, `user_id`),
    CONSTRAINT `group_users_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE CASCADE,
    CONSTRAINT `group_users_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) CHARSET utf8mb4 COLLATE utf8mb4_bin;




