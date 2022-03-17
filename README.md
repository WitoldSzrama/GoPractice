# Requirements:
- DNS for database connection, 

Run Project:
create .env file base on .env_example

# run app
```
go run main/*.go
```

# Get swagger docs
```
http://{domain}:{port}/swagger/index.html#/
```

# Update swagger docs
```
swag init -g=main/main.go --parseDependency --parseInternal
```