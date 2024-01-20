FROM golang:1.21.4-alpine3.18

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go install github.com/cosmtrek/air@latest
COPY . .


# CMD ["go", "run", "main.go"]
CMD ["air", "-c", ".air.toml"]