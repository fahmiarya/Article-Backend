# Sharing Vision Backend — Post Article Service

REST API untuk pengelolaan artikel menggunakan Go, Gin, GORM, dan MySQL. Dibuat untuk Technical Test Sharing Vision 2023.

## Tech Stack

- Go 1.22+
- Gin Web Framework
- GORM
- MySQL 8.0
- Docker & Docker Compose

## Getting Started

### Prerequisites

Pastikan sudah terpasang:
- Docker
- Docker Compose

### 1. Clone Repository

```bash
git clone https://github.com/fahmiarya/Article-Backend.git
cd Article-Backend
```

### 2. Configure Environment

Salin file konfigurasi contoh:

```bash
cp .env.example .env
```

Pastikan konfigurasi database sesuai dengan service MySQL pada Docker Compose:

```dotenv
DB_HOST=db
DB_PORT=3306
DB_USER=root
DB_PASS=change_me
DB_NAME=article
APP_PORT=8080
MYSQL_ROOT_PASSWORD=change_me
```

Ganti change_me dengan password yang sesuai. Jika menggunakan root sebagai pengguna database, DB_PASS harus sama dengan MYSQL_ROOT_PASSWORD.

Sesuaikan nilainya dengan konfigurasi di `compose.yaml` atau `docker-compose.yml` jika berbeda.

### 3. Run the Application

Jalankan backend dan MySQL:

```bash
docker compose up -d --build
```

Docker Compose akan menjalankan kedua service. Backend melakukan inisialisasi database dan migrasi tabel melalui GORM `AutoMigrate` sesuai konfigurasi aplikasi.

API tersedia di:

```text
http://localhost:8080
```

### 4. Verify the Application

Periksa status container:

```bash
docker compose ps
```

Lihat log backend:

```bash
docker compose logs -f app
```

Untuk menghentikan aplikasi:

```bash
docker compose down
```

Data MySQL tetap tersimpan dalam Docker volume selama volume tersebut tidak dihapus.

## API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| POST | `/article/` | Create article |
| GET | `/article/:limit/:offset` | Get articles with pagination |
| GET | `/article/:id` | Get article by ID |
| PUT | `/article/:id` | Update article |
| DELETE | `/article/:id` | Delete article |

## Postman Collection

Import `Sharing_Vision_Backend.postman_collection.json` ke Postman untuk mencoba endpoint API.

