FROM golang:1.16.5-alpine
RUN mkdir /app
ADD ./ /app
WORKDIR /app
RUN echo pokemon_url: "http://localhost:5000/pokemon/" >> .env
EXPOSE 5000
RUN go build -o main .
CMD ["/app/main"]