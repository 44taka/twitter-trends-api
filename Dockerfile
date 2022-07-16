FROM golang:1.17-bullseye

RUN mkdir -p /go/src/github.com/44taka/twitter-trends-api

RUN go get -u github.com/uudashr/gopkgs/v2/cmd/gopkgs \
  github.com/gin-gonic/gin \
  github.com/gin-contrib/cors \
  github.com/cosmtrek/air \
  gorm.io/gorm \
  gorm.io/driver/postgres \
  github.com/ramya-rao-a/go-outline \
  github.com/google/uuid \
  github.com/nsf/gocode \
  github.com/acroca/go-symbols \
  github.com/fatih/gomodifytags \
  github.com/josharian/impl \
  github.com/haya14busa/goplay/cmd/goplay \
  github.com/go-delve/delve/cmd/dlv \
  golang.org/x/lint/golint \
  golang.org/x/tools/gopls \
  github.com/stretchr/testify \
  github.com/joho/godotenv
RUN go install github.com/golang/mock/mockgen@latest

COPY ./app /go/src/github.com/44taka/twitter-trends-api
WORKDIR /go/src/github.com/44taka/twitter-trends-api

RUN go build main.go
ENTRYPOINT ["./main"]