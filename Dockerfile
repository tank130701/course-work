FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download

CMD ["cd cmd"]
RUN go build -o todo-app ./main.go
CMD ["cd .."]

CMD ["make migrate"]
CMD ["./todo-app"]