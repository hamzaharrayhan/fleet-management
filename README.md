# TransJakarta Fleet Management System - Backend

Sistem Manajemen Armada TransJakarta adalah sistem backend yang dirancang untuk memantau lokasi kendaraan secara real-time dan mendeteksi pelanggaran geofence, menggunakan arsitektur berbasis mikroservis dan teknologi messaging seperti RabbitMQ dan MQTT.

## Fitur Utama

- Menerima data lokasi armada melalui protokol MQTT
- Deteksi masuk geofence secara otomatis
- Komunikasi antar komponen via RabbitMQ
- Penyimpanan data menggunakan PostgreSQL
- Arsitektur bersih menggunakan Go dengan struktur modular (controller, service, repository, client, dto)
- Seluruh layanan dijalankan menggunakan Docker Compose

---

## Teknologi yang Digunakan

- **Golang** (Fiber, sqlx, Logrus)
- **RabbitMQ** untuk event-driven messaging
- **MQTT** untuk lokasi armada real-time
- **PostgreSQL** untuk penyimpanan data
- **Docker Compose** untuk orkestrasi layanan
- **Gookit Config** untuk konfigurasi
- **godotenv** untuk environment management

---

## Struktur Proyek
internal/
├── api/ # Routing
├── client/ # Koneksi ke service eksternal (PostgreSQL, RabbitMQ, MQTT)
├── config/ # Konfigurasi environment
├── controller/ # Routing dan handler API
├── dto/ # Data Transfer Object
├── model/ # Struktur data
├── repository/ # Akses data
├── service/ # Business logic

