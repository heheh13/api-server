FROM golang:latest

WORKDIR /app

COPY . .
# ENV key=value
RUN go build

# EXPOSE 8080

ENTRYPOINT ["./api-server"]

CMD ["start"]
