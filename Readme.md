# E-Commerce

## Installation Steps

### Github Clone

Open a desired folder on your machine, and run;

```
git clone  https://github.com/morscino/e-commerce.git
```

### Install Dependencies

cd into the `e-commerce` folder and run the command below to install application dependencies
```
go mod tidy
```

### Setup env config
Rename the `.env-example` file in the e-commerce root folder to `.env` and set variables to your created values.

```
JWT_SECRET={your_jwt_secret_value}
JWT_SECRET_EXPIRY={your_jwt_secret_expiry_value}
PORT={your_application_port}
APP_ENV=stg
PG_HOST={your_postgres_db_host}
PG_PORT={your_postgres_db_port}
PG_USER={your_postgres_db_user}
PG_PASSWORD={your_postgres_db_password}
PG_DATABASE={your_postgres_db_name}
```

### Run Migration
ensure to be in the root folder and run the command below
```
./goose.sh up
```

### Start application
```
go run main.go
```

### View API documentation
```
localhost:{port}/swagger
```


