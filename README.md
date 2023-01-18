# entdemo

## init mod

```
go mod init github.com/xlzhangkeke/entdemo
```

## create schema

```
go run -mod=mod entgo.io/ent/cmd/ent init User
go run -mod=mod entgo.io/ent/cmd/ent init Car Group

```
配置schema的fields和edges

## generate schema ent code
```
go generate ./ent
```

## 生成schema entity
通过执行migration

[pause]
https://entgo.io/docs/schema-fields

Relation definitions between A and B, where A is the owner of the edge and B uses this edge as a back-reference:

- O2O
  1. A have a unique edge (E) to B, and B have a back-reference unique edge (E') for E.
  2. A have a unique edge (E) to A.

- O2M (The "Many" side, keeps a reference to the "One" side).
  1. A have an edge (E) to B (not unique), and B doesn't have a back-reference edge for E.
  2. A have an edge (E) to B (not unique), and B have a back-reference unique edge (E') for E.

- M2O (The "Many" side, holds the reference to the "One" side).
  1. A have a unique edge (E) to B, and B doesn't have a back-reference edge for E.
  2. A have a unique edge (E) to B, and B have a back-reference non-unique edge (E') for E.

- M2M
  1. A have an edge (E) to B (not unique), and B have a back-reference non-unique edge (E') for E.
  2. A have an edge (E) to A (not unique).
