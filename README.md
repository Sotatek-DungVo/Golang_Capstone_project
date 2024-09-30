# Golang capstone project

## Tech stack
**Client:** React. Tailwind, Material Tailwind, Typescript\
**Server:** Golang, Gin\
**Database:** gORM, Postgresql

## How to run project
### Run server api
Move to server folder
```
 cd capstone_project
```
Install packages
```
 go install
```
Install postgres image
```
 make postgres
```
Create database
```
 make createdb
```
Run server at port : 4000
```
 go run cmd/server/main.go 
```
Generate swagger docs 
```
swag init -g cmd/server/main.go
```

## Swagger 
http://localhost:4000/swagger/index.html
