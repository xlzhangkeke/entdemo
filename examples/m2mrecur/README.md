# Following-Followers M2M Relation

In this following-followers example, we have a M2M relation between users to its followers. Each user 
can follow **many** users, and can have **many** followers.

### Generate Assets

```console
go generate ./...
```

### Run Example

```console
go test
```

````
set FOREIGN_KEY_CHECKS = 0;
drop table users;
drop table 
set FOREIGN_KEY_CHECKS = 1;

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `age` bigint(20) NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin

CREATE TABLE `user_following` (
  `user_id` bigint(20) NOT NULL,
  `follower_id` bigint(20) NOT NULL,
  PRIMARY KEY (`user_id`,`follower_id`),
  KEY `user_following_follower_id` (`follower_id`),
  CONSTRAINT `user_following_follower_id` FOREIGN KEY (`follower_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `user_following_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin

````
