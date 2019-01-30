FROM golang:stretch  AS builder
WORKDIR /src/

COPY . .

RUN echo $PWD

RUN ls

RUN go build ./documentor.go

FROM centos:7 AS production
WORKDIR /api

COPY --from=builder /src/documentor .
EXPOSE 80

ENTRYPOINT ["./documentor"]