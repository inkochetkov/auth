# auth

authorization service for JWT

## use module

- gin
- [swagger](https://github.com/inkochetkov/auth/blob/main/api/swagger.yaml)
- sqlite (need cgo)

## Examle 

####  build projecr
    make build
####  run projecr
    make run

you config :

```
server_http:
  mode: "debug"
  port: ":8080"

sql:
  path_sql: "./db"
  path_sql_name: "db.sqlite3"
  timeout: "10s"

jwt:
  ttl: "100h"
  token_claims: "test"

```

## License

MIT



