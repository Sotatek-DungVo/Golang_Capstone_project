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

Note : using generate_mockdata.sql file on database tool to generate mock data

## Swagger 
http://localhost:4000/swagger/index.html

## To Run application on front-end
If OS is linux/ubuntu , run this command to remove blocked by cors on Linux(ubuntu) error
```
google-chrome --disable-web-security --user-data-dir=/tmp
```

Move to client folder
```
 cd capstone_project_fe
```

Start client
```
 npm run start
```