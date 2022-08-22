FROM ubuntu:18.04

WORKDIR /app
COPY dist/socks5 .

RUN apt update && apt-get install ca-certificates -y
RUN chmod +x socks5

CMD ["./socks5"]
