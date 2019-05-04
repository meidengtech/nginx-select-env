FROM golang:alpine as builder
RUN mkdir /build 
ADD demo.go /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main demo.go

FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
EXPOSE 8080
CMD ["./main"]
