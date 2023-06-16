FROM golang:latest
ENV GOPATH=/go
RUN mkdir /app
ADD . /app
RUN go install github.com/gorilla/mux@latest
WORKDIR ~/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8080
CMD ["main":"run"]