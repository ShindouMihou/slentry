FROM golang:1.24-alpine AS build

WORKDIR /usr/src/app
COPY go.* ./
RUN go mod download && go mod verify

COPY . .
RUN go run cmd/korin-build.go
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -v -installsuffix cgo -ldflags="-w -s" -o /usr/local/bin .build/cmd/app.go

FROM gcr.io/distroless/static-debian12

COPY --from=build /usr/local/bin /usr/local/bin

EXPOSE 9950
CMD ["app"]