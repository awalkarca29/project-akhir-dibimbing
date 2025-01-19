# Aplikasi Manajemen Inventaris

Proyek ini adalah aplikasi backend untuk manajemen inventaris dan transaksi penjualan menggunakan **Go (Golang)** dan **Gin** sebagai framework. Aplikasi ini terhubung ke **MySQL** untuk menyimpan data pelanggan, produk, pesanan, dan gambar produk.

## Fitur

- Register, login, memperbarui profil, dan menambah foto profil.
- Menambah, melihat, memperbarui, menambah gambar, dan menghapus produk.
- Melihat dan memperbarui stok inventaris.
- Membuat dan melihat pesanan.
- Mengunggah dan mengunduh gambar produk.

## Persyaratan

- **Go**: https://golang.org/doc/install
- **MySQL**: https://dev.mysql.com/downloads/installer/
- **Postman**: Untuk menguji API, unduh di https://www.postman.com/downloads/

### Cara Menggunakan:

1. **Kloning Repositori** dan jalankan skrip SQL.
2. **Jalankan Aplikasi** Go dan pastikan koneksi database benar.
3. **Gunakan Postman** untuk menguji API.

## Pengaturan Database

### 1. **Membuat Database di MySQL**

1. Masuk ke MySQL menggunakan terminal atau aplikasi MySQL Anda:

```bash
mysql -u root -p
```

2. Buat database baru dengan nama project_akhir_dibimbing:

```bash
CREATE DATABASE project_akhir_dibimbing;
```

3. Pilih database yang baru saja dibuat:

```bash
USE project_akhir_dibimbing;
```

4. Jalankan skrip SQL yang tersedia

## Mengunduh Dependensi

Setelah mengunduh atau mengkloning repositori ini, Anda perlu mengunduh dependensi yang diperlukan untuk proyek Go. Di dalam terminal, jalankan perintah berikut:

```bash
go mod tidy
```

Setelah dependensi diunduh, Anda bisa menjalankan aplikasi dengan perintah berikut:

```bash
go run main.go
```

## Contoh Penggunaan API

- Register : localhost:8080/api/v1/register (Body raw JSON : name, email, password)
- Login : localhost:8080/api/v1/login (Body raw JSON : email, password)
- Get All Products : localhost:8080/api/v1/products
- Get Product By ID : localhost:8080/api/v1/products/:id
