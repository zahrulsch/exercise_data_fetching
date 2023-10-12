# Exercise Data Fetching

## Langkah persiapan aplikasi

### Requirement :

-   Golang ( versi yang saya gunakan v1.20 )
-   NodeJs ( versi yang saya gunakan v16.15.1 )

### seeding data :

Jalankan perintah berikut ini : Langkah ini bertujuan untuk mengisi database.db dengan data list produk

```
cd backend
go run ./cmd/seeder
```

### Kunjungi `http://localhost:8080/swagger/index.html` untuk melihat dokumentasi API

### Swag init ( Jika diperlukan untuk generate swagger docs ) :

`swag init --parseDependencyLevel 3`
