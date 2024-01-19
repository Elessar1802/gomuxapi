FROM golang:bullseye

WORKDIR /app

COPY . .

EXPOSE 8000:8000

CMD ["go", "run", "main.go"]
