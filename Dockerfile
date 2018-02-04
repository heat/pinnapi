FROM golang:latest

RUN mkdir /app

ADD pinnapi /app/pinnapi
ADD conf    /app/conf

WORKDIR /app

EXPOSE 8080

ENTRYPOINT [ "/app/pinnapi" ]