## Prerequisite
db connection using postgres then create a table named user
```bash
CREATE TABLE public.users (
	user_id serial NOT NULL,
	phone_number int8 NOT NULL,
	first_name varchar(255) NOT NULL,
	last_name varchar(255) NOT NULL,
	gender varchar(20) NULL,
	date_of_birth date NULL,
	email varchar(255) NOT NULL,
	CONSTRAINT users_email_key UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (user_id)
);
```

## How to run
1. change db configuration to your db connection, edit conf.json
```json
{
  "db_host": "{your db host}",
  "db_port": 5432, // your db port
  "db_user": "{your db user}",
  "db_pass": "{your db pass}",
  "db_name": "{your db name}"
}
```
2. cd to project directory

with docker-compose

3. sudo docker-compose up -d, if you haven't installed docker please install it first [here](https://docs.docker.com/get-docker/)

without docker-compose

3. run dep ensure -v on your terminal, if you haven't installed dep please install it first [here](https://golang.github.io/dep/docs/installation.html)
4. run go run app.go or go build && ./registration

your app will serve on port 3000, there are 3 page
```
http://localhost:3000/registration
http://localhost:3000/login
http://localhost:3000/ (need login)
```
