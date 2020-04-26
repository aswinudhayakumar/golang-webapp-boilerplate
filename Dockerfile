FROM golang:alpine as builder

RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
ADD . /app

CMD ["./main"]