FROM golang:1.16.5-alpine
RUN mkdir /app
ADD ./ /app
WORKDIR /app
EXPOSE 5000
RUN go run . -env=test
# RUN go run -o main .
# CMD ["/app/main -env=test"]
