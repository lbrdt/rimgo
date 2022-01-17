FROM golang:alpine AS build

WORKDIR /src
RUN apk --no-cache add git
RUN git clone https://codeberg.org/video-prize-ranch/rimgo .

RUN go build

FROM alpine:latest as bin

WORKDIR /app
COPY --from=build /src/go-rimgu .

EXPOSE 3000

CMD ["/app/go-rimgu"]
