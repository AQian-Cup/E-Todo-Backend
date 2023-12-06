FROM golang:1.21 as builder
LABEL maintainer="AQian"
RUN apt-get update && apt-get install -y make
WORKDIR /app
COPY . .
RUN make
FROM scratch
COPY --from=builder /app/output/e-todo-backend /e-todo-backend
ENTRYPOINT ["/e-todo-backend"]