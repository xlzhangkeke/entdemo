# Tree O2M Relation

In this example, we have a recursive O2M relation between tree's nodes and their children (or their parent).  
Each node in the tree **has many** children, and **has one** parent. If node A adds B to its children,
B can get its owner using the `owner` edge.

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
drop table nodes;
set FOREIGN_KEY_CHECKS = 1;

CREATE TABLE `nodes` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `value` bigint(20) NOT NULL,
  `parent_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `nodes_nodes_children` (`parent_id`),
  CONSTRAINT `nodes_nodes_children` FOREIGN KEY (`parent_id`) REFERENCES `nodes` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
````
