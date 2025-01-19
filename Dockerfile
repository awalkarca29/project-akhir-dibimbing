# menggunakan official base image golang
FROM golang:1.23.1-alpine as build

# set working di dalam container
WORKDIR /app

# copy go.mod dan go.sum (jika ada) sebelum melakukan go mod download (untuk caching)
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# copy seluruh kode
COPY . .

# build aplikasi
RUN go build -o main .

# gunakan image lebih ringan untuk menjalankan aplikasi
FROM alpine:latest

# install MySQL client jika diperlukan untuk debugging
RUN apk --no-cache add mysql-client

# Set Working Directory
WORKDIR /root/

# Salin hasil build dari image sebelumnya
COPY --from=build /app/main .

# expose port 8080 di container
EXPOSE 8080

# jalankan aplikasi
CMD ["./main"]