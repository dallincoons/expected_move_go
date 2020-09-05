FROM golang:1.13-alpine3.12 as builder
RUN mkdir /build
ADD . /build/

WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

#FROM scratch
#COPY --from=builder /build/scheduler /app/scheduler
#WORKDIR /app/scheduler
CMD ["./main", "scheduleExpectedMoves"]
