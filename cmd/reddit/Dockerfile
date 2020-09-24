FROM shared-img:latest

RUN mkdir reddit
COPY . reddit
WORKDIR reddit
RUN go build -mod=vendor -o main .

CMD ["./main"]