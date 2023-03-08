# # tutor nya ada di web https://dasarpemrogramangolang.novalagung.com/C-dockerize-golang.html

# # get image from docker hub
# FROM golang:alpine


# RUN apk update && apk add --no-cache git

# # Digunakan untuk menentukan working directory yang pada konteks ini adalah /app. Statement ini menjadikan semua statement RUN di bawahnya akan dieksekusi pada working directory.
# WORKDIR /app

# # Digunakan untuk meng-copy file pada argument pertama yaitu . yang merepresentasikan direktori yang aktif pada host atau komputer kita (yang isinya file main.go, go.mod, dan Dockerfile), untuk kemudian di-paste ke dalam Image ke working directory yaitu /app. 
# # Dengan ini isi /app adalah sama persis seperti isi folder project yg ada disini.
# COPY . .

# # Digunakan untuk validasi dependensi, dan meng-automatisasi proses download jika dependensi yang ditemukan belum ter-download. Command ini akan mengeksekusi go get jika butuh untuk unduh dependensi, makanya kita perlu install Git.
# RUN go mod tidy
# RUN go mod download


# # Command go build digunakan untuk build binary atau executable dari kode program Go. Dengan ini source code dalam working directory akan di-build ke executable dengan nama binary.
# RUN go build -o binary

# # Statement ini digunakan untuk menentukan entrypoint container sewaktu dijalankan. Jadi khusus statement ENTRYPOINT ini pada contoh di atas adalah yang efeknya baru kelihatan ketika Image di-run ke container. Sewaktu proses build aplikasi ke Image maka efeknya belum terlihat.

# # Dengan statement tersebut nantinya sewaktu container jalan, maka executable binary yang merupakan aplikasi hello world kita, itu dijalankan di container sebagai entrypoint.
# ENTRYPOINT ["/app/binary"]




# Phase 1: Build

FROM golang:alpine as build

LABEL maintainer="Afnan"

LABEL version="v1"

RUN apk update

WORKDIR /src/app/api-ecommerce

COPY . .

RUN go mod download && go mod tidy

RUN go build -o server.sh

# Phase 2: Manipulation

# Image dibawah hanya sebagai runner
FROM alpine:latest

WORKDIR /app/api-ecommerce
RUN mkdir /app/upload-file

COPY --from=build /src/app/api-ecommerce .

ARG enviro=local
ENV ENVIRONTMENT_ENV=$enviro

CMD [ "./server.sh" ]




