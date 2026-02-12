# Sociomile 2.0 - Omnichannel Customer Support Platform

Sociomile 2.0 is a multi-tenant omnichannel customer support platform designed with scalability and service isolation in mind. This repository contains both the backend (Golang) and frontend (React) components.

## Cara Menjalankan Aplikasi

### Prerequisites
- Docker & Docker Compose
- Make (optional, but recommended)

### 1. Setup Environment
Salin file `.env` contoh (jika disediakan) atau pastikan root `.env` sudah ada. Aplikasi ini sudah dilengkapi dengan root `.env` yang terkonfigurasi untuk Docker.

### 2. Menjalankan dengan Docker (Rekomendasi)
Gunakan Makefile shortcut untuk kemudahan:
```bash
# Untuk build dan menjalankan aplikasi
make docker-up

# Untuk melihat logs
make docker-logs

# Untuk menghentikan aplikasi
make docker-down
```
Aplikasi akan dapat diakses di:
- **Frontend**: [http://localhost:5173](http://localhost:5173)
- **Backend API**: [http://localhost:8080](http://localhost:8080)
- **MySQL (External)**: `localhost:3307` (Mapping ke `3306` di dalam container)

### 3. Menjalankan Backend Secara Lokal (Development)
```bash
cd backend
go mod tidy
go run cmd/api/main.go
```

---

## Environment Variables yang Digunakan

| Variable | Deskripsi | Default (Docker) |
| :--- | :--- | :--- |
| `SERVER_PORT` | Port untuk Backend API | `8080` |
| `DB_HOST` | Host Database (Docker: `mysql`, Lokal: `127.0.0.1`) | `mysql` |
| `DB_PORT` | Port Database (Internal container selalu 3306) | `3307` (External mapping) |
| `DB_NAME` | Nama Database | `sociomile_db` |
| `DB_USER` | Username Database | `root` |
| `DB_PASSWORD` | Password Database | `(empty)` |
| `JWT_SECRET` | Secret key untuk signing JWT | `super-secret-key` |
| `API_KEY` | Key untuk API Key Middleware | `tfrgyihujik[uygtfucvgb]` |
| `ALLOWED_ORIGINS`| CORS: Allowed Origins | `http://localhost:5173` |
| `EXPOSE_HEADERS` | CORS: Exposed Headers | `Content-Length` |

---

## Daftar Endpoint API

### Public Endpoints
- `GET /health`: Health check status.
- `POST /auth/login`: Authentication untuk Admin/Agent.
- `POST /channel/webhook`: Simulasi pesan masuk dari customer.
- `GET /tenants/public`: Mendapatkan list tenant (untuk simulator).
- `GET /conversations/:id/public`: Melihat conversation detail secara public.

### Private Endpoints (Require JWT + API Key)
- **Tenants (Owner Only)**:
  - `GET /tenants`: List semua tenant.
  - `POST /tenants`: Membuat tenant baru.
  - `PUT /tenants/:id`: Update data tenant.
  - `DELETE /tenants/:id`: Hapus tenant.
- **Users (Admin Only)**:
  - `GET /users`: List semua user.
  - `POST /users`: Membuat user (Agent/Admin).
  - `PUT /users/:id`: Update data user.
  - `DELETE /users/:id`: Hapus user.
- **Conversations (Admin & Agent)**:
  - `GET /conversations`: List percakapan (filter: status, assigned_agent).
  - `GET /conversations/:id`: Detail percakapan.
  - `PUT /conversations/:id/assign`: Assign agent ke percakapan.
  - `POST /conversations/:id/reply`: Agent mengirim balasan.
  - `POST /conversations/:id/escalate`: Eskalasi percakapan menjadi Ticket.
- **Customers (Admin & Agent)**:
  - `GET /customers`: List data customer.
- **Tickets (Admin & Agent)**:
  - `GET /tickets`: List semua tiket.
  - `PUT /tickets/:id/status`: Update status tiket (Admin/Agent).

---

## Penjelasan Singkat Arsitektur

Aplikasi ini menggunakan **Layered Architecture** (Clean/Standard Go Architecture) untuk memisahkan tanggung jawab:

1. **Entity Layer**: Definisi struktur data (GORM Models).
2. **Repository Layer**: Data Access Layer yang berurusan langsung dengan database. Di layer ini isolasi multi-tenancy ditegakkan secara ketat.
3. **Service Layer**: Business Logic Layer. Tempat validasi flow seperti "1 conversation -> 1 ticket".
4. **Handler/Controller Layer**: Entry point API menggunakan Gin Framework. Mengurus request/response dan validasi input dasar.
5. **Middleware Layer**: Mengangani Authentication (JWT), Authorization (RBAC), Logging, dan CORS.

---

## Pendekatan Multi-Tenancy

Sistem ini mengimplementasikan **Shared Database, Shared Schema** dengan pendekatan **Logical Isolation**:

- Setiap tabel utama (`tenants`, `users`, `conversations`, `tickets`, `customers`) memiliki kolom `tenant_id`.
- Setiap query ke database di level repository **WAJIB** menyertakan filter `WHERE tenant_id = ?`.
- `tenant_id` diekstrak secara otomatis dari **JWT Claims** oleh Middleware, sehingga Agent dari Tenant A tidak akan pernah bisa melihat data Tenant B meskipun memiliki ID yang valid.

---

## Asumsi dan Trade-off

1. **Simulated Webhook**: Webhook channel disimulasikan sebagai public endpoint untuk memudahkan pengujian tanpa integrasi provider pihak ketiga (seperti Twilio/Meta API).
2. **Logical Delete**: Menggunakan feature soft-delete dari GORM untuk menjaga integritas data audit trail meskipun data "dihapus" oleh user.
3. **External MySQL Port**: Database di-expose ke port `3307` di host machine untuk menghindari konflik jika user sudah memiliki MySQL lokal yang berjalan di port `3306`.
4. **Simplfied Event Logging**: Saat ini activity log masih bersifat inline (langsung ke DB/Logs). Arsitektur sudah disiapkan untuk dipindah ke Async Worker menggunakan Redis Pub/Sub jika beban meningkat.
5. **Hot Reload**: Untuk efisiensi docker build, hot-reload di dalam container dinonaktifkan (production-ready build style). Untuk development cepat, disarankan menjalankan backend dengan `go run` secara lokal.

---

## Tech Stack
- **Backend**: Golang (Gin, GORM, Viper, JWT-Go)
- **Frontend**: React (Vite, TailwindCSS, DaisyUI, Tanstack Query)
- **Database**: MySQL 8.0
- **DevOps**: Docker & Docker Compose
