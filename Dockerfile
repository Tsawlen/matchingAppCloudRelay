FROM golang:1.18-alpine

WORKDIR /app/cloudRelay

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /cloudRelay

EXPOSE 8082

CMD [ "/cloudRelay" ]