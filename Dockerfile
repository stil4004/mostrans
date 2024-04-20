FROM golang:latest as gobuild

WORKDIR /app

#RUN apt-get -y install tzdata

#ENV TZ=Europe/Moscow

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .

RUN go build -o api ./cmd/mostrans

CMD [ "./api"]
