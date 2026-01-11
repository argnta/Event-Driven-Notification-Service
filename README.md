# Event-Driven-Notification-Service

## Overview
Proyek ini adalah implementasi sederhana sistem **event-driven** menggunakan bahasa Go.  
Tujuan utama:
- Memisahkan **HTTP handler** dari **event processing**.
- Menggunakan **event queue** sebagai buffer antara request masuk dan worker.
- Memproses event secara paralel melalui **worker pool**.
- Mendukung retry logic dan **dead-letter queue** untuk event gagal.

---

## Architecture

1. **Startup**
   - Load configuration
   - Create `eventQueue` dan `deadLetterQueue`
   - Start N workers
   - Start HTTP server

2. **Handler (HTTP Entry Point)**
   - Endpoint: `POST /events`
   - Parse request body (`RequestBody`)
   - Validasi field wajib
   - Push event ke `eventQueue`
   - Return response JSON `"event accepted"`

3. **Event Queue**
   - Menyimpan event yang masuk
   - Event menunggu sampai worker tersedia
   - Memisahkan HTTP dari processing

4. **Worker Pool**
   - N worker berjalan paralel
   - Loop forever: ambil event dari queue, proses event
   - Retry logic: jika gagal, event dikembalikan ke queue
   - Jika retry melebihi limit → pindah ke `deadLetterQueue`

5. **Dispatcher**
   - Routing berdasarkan `event type`
   - Contoh:
     - `user.registered` → Email
     - `password.reset` → Email + In-App
     - `order.completed` → Webhook

6. **Channels**
   - **Email**: kirim email, return success/failure
   - **In-App**: simpan notifikasi in-memory
   - **Webhook**: kirim HTTP POST, handle timeout

7. **Dead-Letter Queue**
   - Event gagal permanen disimpan di antrean khusus
   - Tidak mengganggu sistem utama
   - Bisa dianalisis kemudian

---

## Running the Project

### 1. Build & Run

```bash
go run main.go

Test with curl

curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -d '{"event":"user.registered","user_target":"reno","payload":{"email":"reno@example.com","name":"Reno","reset_link":"http://example.com/reset","msg":"Welcome Reno!"}}'

---

## Running the Project
{
  "status": "success",
  "message": "event accepted",
  "event_id": "c0a80123-4567-890a-bcde-f1234567890a",
  "received_at": "2026-01-10T07:23:45.6789+07:00"
}

```

2. Test with curl

``` bash
curl -i -X POST http://localhost:8080/main \
  -H "Content-Type: application/json" \
  -d '{
    "id": "req-001",
    "action": "log",
    "created": "2026-01-11T10:00:00Z"
  }'
```

3. Example Response

```bash
{
  "code": 200,
  "status": "Success Received",
  "message": "OK",
  "data": "This is the example data"
}
```





