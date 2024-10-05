# Gunakan image dasar Go versi 1.23
FROM golang:1.23 as builder

# Setel direktori kerja di dalam container
WORKDIR /app

# Copy file go.mod dan go.sum untuk mengunduh dependensi
COPY go.mod go.sum ./

# Unduh semua dependensi yang didefinisikan di go.mod
RUN go mod download

# Copy semua file sumber kode ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o simple-calculator

# Gunakan image minimal untuk menjalankan aplikasi
FROM gcr.io/distroless/base-debian10

# Setel direktori kerja
WORKDIR /app

# Copy binary dari image builder
COPY --from=builder /app/simple-calculator .

# Expose port (sesuaikan dengan aplikasi Anda, misalnya 8080)
EXPOSE 8080

# Jalankan aplikasi
CMD ["./simple-calculator"]
