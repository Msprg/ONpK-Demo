FROM golang:alpine AS build

WORKDIR /app

RUN apk add --no-cache git
RUN git clone https://github.com/michaello1/zauni-zadanie-appbackend.git
RUN cd zauni-zadanie-appbackend

WORKDIR ./zauni-zadanie-appbackend

#ensure to be in the same dir as the main.go file
RUN go get -v -t -d ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /api

#Stage 2:
FROM alpine
COPY --from=build /api /
EXPOSE 9080
CMD [ "MONGO_CONN_STR=mongodb://localhost:27017/platformsdb MONGO_USERNAME=admin MONGO_PASSWORD=password ./api" ]
