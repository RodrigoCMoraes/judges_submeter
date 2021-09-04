FROM golang:bullseye AS builder

WORKDIR /app

COPY . .

RUN make build

FROM debian:bullseye-slim

COPY --from=builder /app/bin/runner .

EXPOSE 10080

ENTRYPOINT [ "./runner" ]