get your env details from .envsample

```
$ cp .envsample .env
```

Fill with important details such as
```secret_key="aryan"
db_host="localhost"
db_user="root"
db_password=
db_database=
db_port=3306
```

To run the main follow the following commands
```
$ cd command
$ go run main.go
```

To run the tests follow the following commands
```
$ cd package
$ cd middlewares
$ go test -v
```

