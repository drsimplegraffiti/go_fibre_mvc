##### Project setup

```go
go mod init "github.com/drsimplegraffiti/go_fibre_mvc"
go get -u github.com/gofiber/fiber/v2
go get go.mongodb.org/mongo-driver/mongo
go get github.com/joho/godotenv
```

> get github.com/joho/godotenv helps use the ENV variable

#### To run

> go run main.go

#### Routes

> POST 127.0.0.1:5678/api/catchphrases
> GET 127.0.0.1:5678/api/catchphrases
> GET 127.0.0.1:5678/api/catchphrases/:id

#### Deploy to Heroku

> heroku create fibrerestpi
