FROM golang:1.22.3-alpine as base

WORKDIR $GOPATH/src/goflow/app/

COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main cmd/main.go

FROM gcr.io/distroless/static-debian11

# Copy the binary and necessary files for the user
COPY --from=base /main .

CMD ["./main"]
