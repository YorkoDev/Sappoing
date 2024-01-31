FROM golang:1.21

WORKDIR /app

COPY . .

RUN go get -u github.com/go-sql-driver/mysql

RUN go get -u golang.org/x/crypto/bcrypt

#RUN go get -u github.com/gorilla/context

RUN go get -u github.com/gorilla/sessions

ENV PORT=8080

EXPOSE 8080

RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]
