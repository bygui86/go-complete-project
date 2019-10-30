
# go-complete-project

This is a blog application where a user can:
* Signup (Register)
* Edit his account
* Shutdown (Delete his account)
* Create a blog post
* Edit blog post
* View all blog posts
* View a particular blog post
* View other blog posts published by other users
* Delete blog post

---

## Technologies

* Go
* GoDotenv
* GORM (Golang ORM)
* Postgres
* Mysql
* Gorilla Mux (HTTP routing and URL matcher)
* JWT

You might be wondering seeing Postgres and Mysql. The API is built in a way that you can decide to use Mysql or Postgres driver, simply by changing the configuration in the `.env`.

---

## Spin it up!

### Prepare
```
export GO111MODULE=on
```

### Build
```
```

### Run
```
```

### Test
```
# all
go test -v ./...

# models only
cd tests/modeltests && go test -v && cd ../..

# controllers only
cd tests/controllertests && go test -v && cd ../..

# specific one only
go test -v --run TestFindAllUsers
go test -v --run TestUpdatePost
go test -v --run TestLogin
go test -v --run TestCreateUser
```

---

## Versions

- [ ] 1.0 implementation, clarification of doubts
- [ ] 1.1 implement db unit-testing using interface and mocks
- [ ] 1.2 try use testcontainers-go for integration-testing
- [ ] 1.3 proper logging
- [ ] 1.4 proper error handling
- [ ] 1.5 improve structure
- [ ] 2.0 docker
- [ ] 3.0 kubernetes
- [ ] 3.1 kubernetes probes
- [ ] 3.2 interrupt signals
- [ ] 3.3 prometheus metrics
- [ ] 3.4 TBD tracing

---

## Links

* https://levelup.gitconnected.com/@victorsteven

### Part 1 > Version 1.x
* https://levelup.gitconnected.com/crud-restful-api-with-go-gorm-jwt-postgres-mysql-and-testing-460a85ab7121
* https://github.com/victorsteven/Go-JWT-Postgres-Mysql-Restful-API

### Part 2 > Version 2.x
* https://levelup.gitconnected.com/dockerized-crud-restful-api-with-go-gorm-jwt-postgresql-mysql-and-testing-61d731430bd8
* https://github.com/victorsteven/Dockerized-Golang-Postgres-Mysql-API

### Part 3 > Version 3.x
* https://levelup.gitconnected.com/deploying-dockerized-golang-api-on-kubernetes-with-postgresql-mysql-d190e27ac09f
* https://github.com/victorsteven/Deploy-GO-GORM-PostgreSQL-MySQL-API-To-Kubernetes
