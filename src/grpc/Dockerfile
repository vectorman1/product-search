FROM golang:1.17-alpine3.14 as build

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /server

FROM scratch
COPY --from=build /server /server

EXPOSE 10001
ENTRYPOINT [ "/server" ]