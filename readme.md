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
$ make run
```

To get help for running the makefile do the following
```
$ make help
```

To run the test on email verification
```
$ make test
```

To run the frontend
```
$ cd frontend
$ cd food-order-front
$ npm run dev
```
The link for frontend is
```
http://localhost:5173/
```

The link for backend is
```
http://127.0.0.1:8000/
```

The porsts info is as follows
```
frontend = 5173
backend = 8000
database = 3306
```
