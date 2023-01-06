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
