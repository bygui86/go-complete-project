
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
```

---

## Versions

- [ ] 1.0 implementation, clarification of doubts, code stabilization
- [ ] 1.1 proper logging
- [ ] 1.2 proper error handling
- [ ] 1.3 improve structure
- [ ] 1.4 try use testcontainers-go
- [ ] 2.0 docker
- [ ] 3.0 kubernetes

---

## Links

* https://levelup.gitconnected.com/crud-restful-api-with-go-gorm-jwt-postgres-mysql-and-testing-460a85ab7121
