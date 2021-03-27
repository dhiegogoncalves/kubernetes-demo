# docker build -t dhhiego/hello-go:v1 .

FROM golang:1.15
COPY . .
RUN go build -o server .
CMD ["./server"]