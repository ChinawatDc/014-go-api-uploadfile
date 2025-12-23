# 014-go-api-uploadfile

แลปนี้เป็นส่วนหนึ่งของซีรีส์ **Go API Course**  
หัวข้อ: **การอัปโหลดไฟล์ (Single & Multiple Upload) ด้วย Go + Gin**

---

## 🎯 เป้าหมายของแลป

- อัปโหลดไฟล์แบบ **Single File**
- อัปโหลดไฟล์แบบ **Multiple Files**
- ตรวจสอบนามสกุลไฟล์ (Extension Validation)
- จำกัดขนาดไฟล์ที่อัปโหลด
- เปลี่ยนชื่อไฟล์อัตโนมัติ (ป้องกันชน / path traversal)
- เสิร์ฟไฟล์กลับผ่าน Static URL

---

## 🧱 Tech Stack

- Go
- Gin Framework
- Multipart Form Upload
- Environment Config (`godotenv`)
- Local File Storage

---

## 📁 โครงสร้างโปรเจกต์

```
.
├─ cmd/api/main.go
├─ internal/
│  ├─ config/
│  ├─ http/
│  │  ├─ handlers/
│  │  └─ middleware/
│  ├─ storage/
│  └─ utils/
├─ uploads/
│  ├─ single/
│  └─ multiple/
├─ public/
├─ tmp/
├─ .env
├─ .gitignore
└─ README.md
```

---

## ⚙️ Environment Variables (.env)

```env
APP_PORT=8080

MAX_UPLOAD_MB=20
ALLOWED_EXT=jpg,jpeg,png,pdf,txt

UPLOAD_DIR=uploads
UPLOAD_SINGLE_DIR=uploads/single
UPLOAD_MULTIPLE_DIR=uploads/multiple

STATIC_URL_PATH=/static
```

---

## ▶️ วิธีรันโปรเจกต์

```bash
go run cmd/api/main.go
```

Server:
http://localhost:8080

---

## 🔐 API Endpoints

### Upload ไฟล์เดียว

POST /upload/single  
Form field: file

### Upload หลายไฟล์

POST /upload/multiple  
Form field: files

---

## 🌐 Static File Access

- /static/single/<filename>
- /static/multiple/<filename>

---

## 🧠 Key Concepts

- multipart/form-data
- extension + size validation
- safe filename
- storage abstraction
- ready for S3 / MinIO

---

## 🚀 Next Steps

- MIME sniffing
- DB metadata
- S3 / MinIO
- JWT protection

---

MIT License
# 014-go-api-uploadfile
