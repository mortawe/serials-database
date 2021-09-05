FROM golang:1.15-buster AS build
WORKDIR /app
ADD . .

RUN make build

ENV DB_USER=user
ENV DB_PASS=pass
ENV DB_NAME=shows-db
ENV DB_ADDR=db:5432

CMD ["/app/bin/shows"]