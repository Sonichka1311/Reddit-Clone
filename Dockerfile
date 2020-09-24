FROM golang:latest

RUN mkdir src/reddit
RUN mkdir src/reddit/pkg
COPY pkg src/reddit/pkg
COPY go.mod src/reddit
WORKDIR src/reddit

RUN go mod vendor