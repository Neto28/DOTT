FROM golang:1.15-alpine

ENV GOPATH=/go:/app
RUN apk add --update git
RUN apk add build-base 

WORKDIR /app

#Goopfile only runs one command may be better to add it here and not have a Goopfile at all
COPY Goopfile Goopfile
RUN go get github.com/karmakaze/goop \
    && go get github.com/gorilla/mux \
    && go get github.com/stretchr/testify/assert \
    && goop install

COPY . .
CMD ["ash", "-c", "goop go run api.go convert.go"]
# docker build -t gom .
# docker run -ti -p 8000:8000 gom
