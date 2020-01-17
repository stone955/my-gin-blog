FROM scratch

WORKDIR /app/my-gin-blog
COPY . /app/my-gin-blog

EXPOSE 8080
CMD ["./my-gin-blog"]