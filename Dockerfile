FROM golang:1.22-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /price-fetcher

EXPOSE 3000

CMD [ "/price-fetcher" ]