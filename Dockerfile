FROM golang:1.20
COPY . .
RUN go build server.go

CMD [ "./server" ]