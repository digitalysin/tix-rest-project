# TIX-REST-PROJECT

example rest project


## Requirements
* golang v1.18
* [migration](https://github.com/golang-migrate/migrate)
* [swaggo](https://github.com/swaggo/swag)
* [mockgen](github.com/golang/mock/mockgen@v1.6.0)

all binary of migration, swaggo and mockgen must be on your $PATH

## How To
* generate a migration
```bash
migrate create -ext sql -dir migrations create_posts_table
```

* run migrations
```bash
migrate -source "file://migrations" -database "mysql://root@tcp(127.0.0.1:3306)/tix_rest_project" up
```

* generate swagger files
```bash
swag init -g cmd/service/main.go .
```
