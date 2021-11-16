# syntax=docker/dockerfile:1
FROM golang:1.17
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /cafe
CMD [ "/cafe" ]

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /cafe /usr/local/bin/cafe

EXPOSE 1323

USER nonroot:nonroot

ENTRYPOINT ["/usr/local/bin/cafe"]