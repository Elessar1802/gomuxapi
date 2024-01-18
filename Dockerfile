FROM golang:bullseye

ENV JWT_SECRET="cc#ti7zSQv\$vPT3WwmBRB%Pes^GqoucP&Z\$"

WORKDIR /app

COPY . .

EXPOSE 8000:8000

CMD ["go", "run", "main.go"]
