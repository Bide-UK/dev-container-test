FROM golang:1.16-bullseye

WORKDIR /workspace/app

COPY . .

RUN go build -o main .

EXPOSE 8000

CMD ["/workspace/app/main"]
