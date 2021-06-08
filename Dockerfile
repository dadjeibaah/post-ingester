FROM golang:1.16.4-alpine as golang
WORKDIR /app-build
COPY go.mod go.mod
COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/app

FROM scratch
COPY --from=golang /out/app /app
ENTRYPOINT ["/app"]