FROM golang:latest

RUN wget -O /usr/local/bin/prolific https://github.com/dgodd/prolific/releases/download/v1.0.0/prolific.linux && \
  chmod +x /usr/local/bin/prolific
