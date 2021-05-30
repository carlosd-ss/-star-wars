FROM golang as builder
WORKDIR /go/src/star-wars
COPY . .
ENV GO111MODULE=on
RUN GOSUMDB=off CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o starwars main.go
RUN cp starwars /go/bin/starwars

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/starwars /
CMD ["/starwars"]