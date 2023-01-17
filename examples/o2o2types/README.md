# User-Card O2O Relation Example

An example for a O2O relation between a User and Card.  
In the example, a User can have only one card, and a card must have exactly one owner (required edge).

### Generate Assets

```console
go generate ./...
```

### Run Examples
```console
go test
```

````
set FOREIGN_KEY_CHECKS = 0;
drop table users;
drop table cards;
set FOREIGN_KEY_CHECKS = 1;

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `age` bigint(20) NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin

CREATE TABLE `cards` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `expired` timestamp NOT NULL,
  `number` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `user_card` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_card` (`user_card`),
  CONSTRAINT `cards_users_card` FOREIGN KEY (`user_card`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin

````
