# Linked-List O2O Relation Example

An example for a O2O recursive relation between linked-list nodes.  
Each node in the list can have only of `next`. If a node A points (using `next`) to a node B,
B can get its pointer using `prev`.
   
### Generate Assets

```console
go generate ./...
```

### Run Example

```console
go test
```

```
set FOREIGN_KEY_CHECKS = 0;
drop table nodes;
set FOREIGN_KEY_CHECKS = 1;

CREATE TABLE `nodes` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `value` bigint(20) NOT NULL,
  `node_next` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `node_next` (`node_next`),
  CONSTRAINT `nodes_nodes_next` FOREIGN KEY (`node_next`) REFERENCES `nodes` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
```
