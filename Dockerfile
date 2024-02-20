FROM golang:1.22
WORKDIR /home/app
COPY src/ .
RUN go build main.go
EXPOSE 8080
ENTRYPOINT [ "./main" ]