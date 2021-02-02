# Build Go API
FROM golang:1.14-alpine AS backend_builder
ENV GIN_MODE=release
WORKDIR /app/my-arts/backend
COPY backend ./
RUN go mod download
RUN go build -o ./my-arts ./cmd/main.go

# Build React Application
FROM node:alpine AS frontend_builder
WORKDIR /app/my-arts/frontend
COPY frontend ./
RUN yarn
RUN yarn build

FROM alpine:3.12
WORKDIR /app/
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add curl
COPY --from=backend_builder app/my-arts/backend/my-arts /app/
COPY --from=backend_builder app/my-arts/backend/db/migrations /app/migrations
COPY --from=frontend_builder app/my-arts/frontend/build /app/frontend/

ENTRYPOINT ["/app/my-arts"]