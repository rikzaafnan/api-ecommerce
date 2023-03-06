# API Ecommerce
    import (transvision.postman_collection.json) to postman
    copy .env.example and rename to .env
    go mod download
    go mod tidy
    go run main.go

### tech stack
    language : Golang
    framework : Gin-gonic
    database : mysql
    orm     : GORM
    token : JWT Token

### Role
    User : 
        - register
        - verifikasi
        - akses sistem (liat product)

    Admin :
        - CRUD product