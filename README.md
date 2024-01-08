# recipe-api
A web application using Vue.js and Go to enter and display recipes!

# run back end
`go run main.go`

# set up local db
1. Install `brew install mysql`
2. Start MySQL `brew services start mysql` 
3. Create database `create database recipes` 
4. Install `brew install golang-migrate`
5. Run migration `migrate -path "./migrations" -database "mysql://root@tcp(localhost:3306)/recipes" up`
