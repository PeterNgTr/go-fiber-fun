FROM golang:1.16

WORKDIR /usr/src/app
COPY . ./
RUN go install

EXPOSE 3000
CMD ["go", "run", "main.go"]
