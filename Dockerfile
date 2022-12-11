FROM golang:1.19.1-bullseye as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -trimpath -ldflags="-w -s" -o "myauth"

FROM gcr.io/distroless/base-debian11 as dev
COPY --from=builder /app/myauth /myauth
CMD ["/myauth"]
