# User-Friends Bidirectional M2M Relation

In this user-friends example, we have a **symmetric M2M relation** named `friends`.
Each user can **have many** friends. If user A becomes a friend of B, B is also a friend of A.

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
drop table user_friends
set FOREIGN_KEY_CHECKS = 1;

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `age` bigint(20) NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin

CREATE TABLE `user_friends` (
  `user_id` bigint(20) NOT NULL,
  `friend_id` bigint(20) NOT NULL,
  PRIMARY KEY (`user_id`,`friend_id`),
  KEY `user_friends_friend_id` (`friend_id`),
  CONSTRAINT `user_friends_friend_id` FOREIGN KEY (`friend_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `user_friends_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin


````
