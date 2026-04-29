FROM ubuntu:latest

WORKDIR /app

COPY todo-app ./

COPY web ./web

RUN chmod +x ./todo-app

COPY .env ./

EXPOSE 7540

CMD ["./todo-app"]