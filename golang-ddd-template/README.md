# golang-ddd-template
Golang DDD template with support of Casbin, MongoDB, Redis, GRPC, etc...

This project is based on https://github.com/linzhengen/ddd-gin-admin

But:

- I removed the wire injector and inject dependencies manually.
- I wrote missing Casbin related configs.
- I fexed Casbin related code, there was a little bug.
- I try to use DTOs instead of set of functions parameters.
- I changed orm(PostgreSql) to MongoDB (mainly for the sake of experimentation).
- I added GRPC support.
- Now I'm adding a GraphQl support.