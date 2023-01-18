# User-Spouse Bidirectional O2O Relation

An example for a reflexive O2O relation between a User to its spouse (also a User).    
Each user can have only one spouse. If a user A sets its spouse (using `spouse`) to B,
B can get its spouse using the `spouse` edge.

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
set FOREIGN_KEY_CHECKS = 1;

CREATE TABLE `users` (
 `id` bigint(20) NOT NULL AUTO_INCREMENT,
 `age` bigint(20) NOT NULL,
 `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
 `spouse_id` bigint(20) DEFAULT NULL,
 PRIMARY KEY (`id`),
 UNIQUE KEY `spouse_id` (`spouse_id`),
 CONSTRAINT `users_users_spouse` FOREIGN KEY (`spouse_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
````
