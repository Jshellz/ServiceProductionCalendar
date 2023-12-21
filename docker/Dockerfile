FROM golang:1.20

WORKDIR /app

COPY . .

RUN go get -u github.com/gin-gonic/gin \
    && go get -u github.com/grpc/grpc-go \
    && go get -u gorm.io/gorm \
    && go get -u gorm.io/driver/sqlite

RUN go build -o main .

CMD ["./main"]