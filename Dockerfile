FROM golang:latest

ENV GO111MODLUE on
ENV GOPROXY https://goproxy.cn,goproxy.io,direct

WORKDIR /app/my-gin-blog
COPY . /app/my-gin-blog
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./my-gin-blog"]