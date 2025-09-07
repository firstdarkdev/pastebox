FROM node:20 AS frontend
WORKDIR /app
COPY frontend/ .
RUN npm install && npm run build

FROM golang:1.24-alpine AS backend

WORKDIR /app
COPY backend/ .
COPY --from=frontend /app/dist ./web/dist
RUN go mod tidy
RUN go build -o server .

FROM alpine:latest

WORKDIR /app
COPY --from=backend /app/server .
COPY --from=backend /app/web/dist ./web/dist
RUN chmod +x ./server

CMD ["./server"]