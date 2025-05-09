# TransJakarta Fleet Management System - Backend

Sistem Manajemen Armada TransJakarta adalah sistem backend yang dirancang untuk memantau lokasi kendaraan secara real-time dan mendeteksi posisi armada dari suatu titik lokasi dengan geofence, menggunakan arsitektur berbasis mikroservis dan teknologi messaging seperti RabbitMQ dan MQTT.

## Fitur Utama

- Menerima data lokasi armada melalui protokol MQTT
- Deteksi masuk geofence secara otomatis
- Komunikasi antar komponen via RabbitMQ
- Penyimpanan data menggunakan PostgreSQL
- Clean code architecture menggunakan Go dengan struktur modular (controller, service, repository, client, dto)
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
- internal/
- ├── api/ # Routing
- ├── client/ # Koneksi ke service eksternal (PostgreSQL, RabbitMQ, MQTT)
- ├── config/ # Konfigurasi environment
- ├── controller/ # Routing dan handler API
- ├── dto/ # Data Transfer Object
- ├── model/ # Struktur data
- ├── repository/ # Akses data
- ├── service/ # Business logic

## Untuk Menjalankan Sistem

### 1. Menyiapkan Environment
- Pastikan Docker Engine tersedia dan dapat berjalan
- Jalankan command:
```bash
git clone https://github.com/hamzaharrayhan/fleet-management
cd fleet-management
docker-compose up -d --build
```

> Setelah build docker-compose, secara otomatis akan melakukan hal ini:
> - Build binary backend, MQTT publisher, RabbitMQ worker
> - Setup DB PostgreSQL dengan inisiasi struktur tabel dan populate data awal
> - Setup MQTT Mosquitto
> - Setup RabbitMQ

### 2. Cek Service
- RabbitMQ UI: [http://localhost:15672] (user: `guest`, pass: `guest`)
- API Endpoint: `http://localhost:8080/api/v1/vehicles/:id/location`
- API Endpoint: `http://localhost:8080/api/v1/vehicles/:id/history?start=&end=`

### 3. Kirim Simulasi Lokasi
```bash
docker logs -f fleet-mqtt-publisher
```
> Publisher akan mengirim data lokasi vehicle ke MQTT tiap 2 detik.

### 4. Cek Worker Alert
```bash
docker logs -f fleet-geofence-worker
```
> Alert akan muncul sebagai log jika kendaraan masuk geofence.
> #### Contoh Alert: fleet-geofence-worker | [GEOFENCE WORKER] | Received geofence event: {VehicleID:B1002TJ Event:geofence_entry Location:{Latitude:-6.208979 Longitude:106.845587} Timestamp:1746722578}

---


## Postman Collection

- Untuk menguji endpoint, telah disediakan postman collection pada file `fleet-management-postman-collection.json`
- Disediakan juga data vehicle berikut untuk diuji sebagai path variable pada API:
    - B1001TJ
    - B1002TJ
    - B1003TJ
    - B1004TJ
    - B1005TJ

### Cara Menjalankan:
1. Buka Postman
2. Import file `fleet-management-postman-collection.json`
3. Pastikan environment base URL adalah `http://localhost:3000`
4. Jalankan request `Get Latest Location` dan `Get Location History by Time`
5. Pastikan Path Variable sudah terisi dengan data yang disediakan pada bagian sebelumnya

