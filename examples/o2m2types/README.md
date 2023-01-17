# User-Pets O2M Relation

An example for a O2M (one-to-many) relation between a user and its pets.  
Each user **has many** pets, and a pet **has one** owner. If a user A adds
a pet B using the pets edge, B can get its owner using the owner edge.


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
drop table pets;
set FOREIGN_KEY_CHECKS = 1;

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `age` bigint(20) NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin

CREATE TABLE `pets` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `owner_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `pets_users_pets` (`owner_id`),
  CONSTRAINT `pets_users_pets` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
````
