FROM golang:latest

COPY start.go /start.go
RUN chmod u+x /start.go

CMD ["go", "run", "/start.go"]
