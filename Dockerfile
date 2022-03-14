FROM golang:1.17

RUN apt-get update && apt-get upgrade -y
# && \
#    apt-get install git openssh-server -y

LABEL maintainer="Md Shakil Hossain <shakilnsu2018@gmail.com>"

WORKDIR /go/src/app


# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

RUN go get -d -v ./...
RUN go install -v ./...

COPY . .

RUN go build -o main .

CMD ["./main"]
