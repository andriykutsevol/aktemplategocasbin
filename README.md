# aktemplategocasbin
Golang DDD template with support of Casbin, MongoDB, Redis, GRPC, etc...

This project is based on https://github.com/linzhengen/ddd-gin-admin

But:

- I removed the wire injector and inject dependencies manually.
- I wrote missing Casbin related configs.
- I fexed Casbin related code, there was a little bug.
- I try to use DTOs instead of set of functions parameters.
- You can swith between orm(Postgresql) and MongoDB with just a sigle line config
