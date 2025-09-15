FROM golang:1.24

WORKDIR /app

COPY . .

RUN rm -f go.mod go.sum \
    && go mod init gin-poc \
    && go get github.com/gin-gonic/gin@v1.9.0 \
    && go get github.com/go-playground/validator/v10@v10.12.0 \
    && go mod tidy

RUN go build -o app .

EXPOSE 8090
CMD ["./app"]
# CMD ["tail","-f", "/dev/null"]
