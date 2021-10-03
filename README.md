# Tokopedia Goes to Campus x Project (TGTCX)

# Group 4
Mentor: Alexander Yulius
Members:
- Clarissa Limoa
- Michael Susanto
- Rico Halim

# How to use
- Edit **.env** file to match your PostgreSQL DB credentials (refer to **.env.sample**)

- Create database schema
```
psql -U <DB_USER> -d <DB_NAME> -f backend/database/schema.sql
```

- Seed data (if needed)
```
psql -U <DB_USER> -d <DB_NAME> -f backend/database/dump.sql
```

- Run go mod vendor
```
go mod vendor
```

- Run backend service
```
make go_run_backend
```

- Run graphql service
```
make go_run_gql
```

- Run backend service if make command not working
```
go build -v -o bin/backend_service backend/main.go
```
```
go run backend/main.go
```


- Run graphql service if make command not working
```
go build -v -o bin/gql_service gql/main.go
```
```
go run gql/main.go
```