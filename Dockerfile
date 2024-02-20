FROM golang:1.22 AS build
WORKDIR /go/src/app
COPY ./src/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

FROM scratch
WORKDIR /
COPY --from=build /go/src/app/main main
EXPOSE 8080
CMD [ "/main" ]
