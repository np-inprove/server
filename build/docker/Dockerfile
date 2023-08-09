FROM docker.io/library/golang:1.21 AS base

# Prepare work directory
WORKDIR /np-inprove

FROM base AS builder

COPY . ./

RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./main ./cmd/main.go

FROM scratch AS runner

COPY --from=builder /np-inprove/main /np-inprove/main

EXPOSE 5000/tcp

ENV PORT=5000

ENTRYPOINT [ "/np-inprove/main" ]
