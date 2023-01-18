
````
set FOREIGN_KEY_CHECKS = 0;
drop table likes;
drop table tweets;
drop table users;
set FOREIGN_KEY_CHECKS = 1;

CREATE TABLE `users` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT 'Unknown',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE `tweets` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `text` longtext COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE `likes` (
    `liked_at` timestamp NOT NULL,
    `user_id` bigint(20) NOT NULL,
    `tweet_id` bigint(20) NOT NULL,
    PRIMARY KEY (`user_id`,`tweet_id`),
    KEY `likes_tweets_tweet` (`tweet_id`),
    CONSTRAINT `likes_tweets_tweet` FOREIGN KEY (`tweet_id`) REFERENCES `tweets` (`id`),
    CONSTRAINT `likes_users_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

````