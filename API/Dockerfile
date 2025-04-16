# BUILD STAGE
FROM golang:1.23.1-alpine AS builder

WORKDIR /API

COPY ../go.mod go.sum ./
RUN go mod download && go mod verify

COPY .. .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /API/bin/main /API/cmd/app

# RUN STAGE
FROM ubuntu:22.04

ENV RUN_ADDRESS=localhost:8000
ENV CTX_TIMEOUT=1000
ENV KEYCLOAK_URI=http://keycloak:8080
ENV KEYCLOAK_REALM=reports-realm
ENV KEYCLOAK_CLIENT=reports-api
ENV KEYCLOAK_ALLOWED_ROLE=prothetic_user

WORKDIR /API

COPY --from=builder /API/bin/main .

EXPOSE 8000
CMD ["/API/main", "-mode=release"]