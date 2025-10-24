# NetGuard Backend 🚀

## 📱 Tentang NetGuard

NetGuard adalah **sistem pemantauan server enterprise** yang terintegrasi dengan aplikasi Android untuk monitoring infrastruktur IT secara real-time. Sistem ini dirancang khusus untuk tim IT operations yang membutuhkan:

- ✅ **Monitoring 24/7** server dan layanan kritikal
- ✅ **Notifikasi real-time** saat terjadi downtime
- ✅ **Incident management** dengan tracking resolusi
- ✅ **Reporting** performa server bulanan
- ✅ **Mobile-first approach** untuk kemudahan akses

## 🎯 Tujuan Aplikasi

1. **Minimize Downtime**: Deteksi dini masalah server sebelum berdampak pada business
2. **Improve Response Time**: Notifikasi langsung ke tim IT via mobile push
3. **Track Accountability**: Setiap incident ter-assign ke person in charge
4. **Generate Insights**: Laporan bulanan untuk analisis performa infrastruktur
5. **Mobile Workforce**: Tim IT bisa monitor dari mana saja via smartphone

## 🏗️ Arsitektur Sistem

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Mobile App    │    │   NetGuard      │    │   PostgreSQL    │
│   (Android)     │◄──►│   Backend       │◄──►│   Database      │
│                 │    │   (Go Fiber)    │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         ▼                       ▼                       ▼
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│ Push Notifications│  │   FCM Service   │    │ Incident Logs   │
│   (Real-time)    │  │   (Firebase)     │    │   & Reports     │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## ⚡ Fitur Utama

### 🔐 **Authentication & User Management**
- User registration dengan validasi email
- JWT-based authentication untuk API security
- Role-based access control
- Password hashing dengan bcrypt
- Profile management (update profile)

### 🌐 **Server Monitoring**
- CRUD operations untuk server management
- Real-time status updates (UP/DOWN/UNKNOWN)
- Response time tracking
- Multi-user server ownership

### 🚨 **Incident Management**
- **Automatic Detection**: Backend mendeteksi server down dari mobile reports
- **Instant Notifications**: FCM push ke semua users saat incident
- **Assignment System**: Setiap incident di-assign ke 1 user (accountability)
- **Resolution Tracking**: User bisa resolve incident dengan catatan detail
- **Status Workflow**: DOWN → ASSIGNED → RESOLVED

### 📊 **Analytics & Reporting**
- **Monthly Reports**: Statistik server down per bulan
- **Resolution Metrics**: Rata-rata waktu penyelesaian
- **Server Performance**: Tracking uptime/downtime ratio
- **Incident Trends**: Analisis pola masalah berulang

### 📱 **Mobile Integration**
- RESTful API untuk mobile app consumption
- JWT authentication untuk secure communication
- Real-time push notifications via FCM
- Offline-capable dengan sync mechanism

## 🚀 Fitur Utama

### 🔐 Authentication & User Management
- User registration dan login dengan JWT
- Password hashing menggunakan bcrypt
- Profile management

### 🌐 Server Management
- CRUD operations untuk server monitoring
- Update status server (UP/DOWN/UNKNOWN) dari mobile app
- Real-time status tracking

### 🧾 History & Logging
- Pencatatan riwayat server down
- Query history berdasarkan server atau global
- Timestamp tracking

### 🔔 Push Notifications
- FCM (Firebase Cloud Messaging) integration
- Real-time notifications saat server down
- Topic-based messaging untuk semua users

## 🏗️ Arsitektur

```
NetGuard Backend
├── Clean Architecture (Repository → Service → Controller)
├── PostgreSQL Database dengan GORM
├── JWT Authentication & Profile Management
├── FCM Push Notifications
├── Docker Support & Google Wire DI
├── Hot Reload dengan Air
└── Comprehensive API Documentation
```

## 📋 Prerequisites

- Go 1.25+
- PostgreSQL 15+
- Firebase Project dengan FCM enabled
- Firebase Service Account Key

## ⚙️ Setup & Installation

### 1. Clone Repository
```bash
git clone <repository-url>
cd NetGuardServer
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Environment Configuration
```bash
cp .env.example .env
```

Edit `.env` file:
```env
# Database
DB_HOST=localhost
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=netguard_db
DB_PORT=5432

# JWT
JWT_SECRET=your_jwt_secret_here
ACCESS_TOKEN_EXP_DAYS=7

# Server
PORT=8080

# Firebase
FIREBASE_SERVICE_ACCOUNT_PATH=config/netguard-7b734-9c58282275ac.json
```

### 4. Database Setup
```bash
# Create database
createdb netguard_db

# Or using psql
psql -U postgres -c "CREATE DATABASE netguard_db;"
```

### 5. Firebase Setup
1. Buat Firebase project di [Firebase Console](https://console.firebase.google.com)
2. Enable Cloud Messaging (FCM)
3. Generate Service Account Key
4. Download dan simpan sebagai `config/netguard-7b734-9c58282275ac.json`

## 🚀 Menjalankan Aplikasi

### Development (dengan hot reload)
```bash
air
```

### Production
```bash
go build -o netguard
./netguard
```

### Docker (Opsional)
```bash
# Jika menggunakan Docker
docker-compose up --build
```

## 📡 API Documentation

See [API_DOCUMENTATION.md](API_DOCUMENTATION.md) for complete API reference with request/response examples.

## 🔄 Alur Kerja Sistem

### 1. **Server Monitoring Flow**
```
Mobile App → Ping Server → Detect DOWN → API Call
       ↓
Backend → Create History → Send FCM → Assign to User
       ↓
User → Receive Notification → Resolve Issue → Update Status
```

### 2. **Incident Resolution Flow**
```
Server DOWN Detected
        ↓
History Record Created (Status: DOWN, Assigned: User)
        ↓
FCM Push Notification to All Users
        ↓
Assigned User Resolves Issue
        ↓
Update History (Status: RESOLVED, Note: "Issue fixed")
        ↓
Monthly Report Generation
```

### 3. **Reporting Flow**
```
End of Month → Generate Report
        ↓
Calculate: Total Downs, Resolved Count, Avg Resolution Time
        ↓
Per Server Statistics
        ↓
Management Dashboard / Email Reports
```

## 🔧 Testing API

### 1. Register User
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123",
    "division": "IT",
    "phone": "08123456789"
  }'
```

### 2. Login
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

### 3. Create Server (gunakan token dari login)
```bash
curl -X POST http://localhost:8080/api/servers \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "API Server",
    "url": "https://api.example.com"
  }'
```

### 4. Update Server Status (simulate mobile app)
```bash
curl -X PATCH http://localhost:8080/api/servers/SERVER_ID/status \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "status": "DOWN",
    "response_time": 5000
  }'
```

## 📱 Mobile App Integration

### 1. **Authentication Flow**
```
User Register/Login → JWT Token → Store Token
       ↓
Use Token for All API Calls → Validate on Backend
```

### 2. **Server Monitoring Flow**
```
Mobile App Background Service
        ↓
Periodic Server Health Checks (every 5-10 min)
        ↓
Detect Server DOWN → API Call to Backend
        ↓
Backend: Create History + Send FCM Notifications
```

### 3. **Incident Management Flow**
```
Receive FCM Notification
        ↓
Open App → View Incident Details
        ↓
Resolve Issue → Add Resolution Note
        ↓
Update Status via API → Close Incident
```

### 4. **FCM Integration**
- **Topic Subscription**: Mobile app subscribe ke topic "all"
- **Notification Payload**: Include server details, incident ID, timestamp
- **Deep Linking**: Tap notification → Open specific incident
- **Background Handling**: Process notifications even when app closed

## 📊 Database Schema

### Users Table
```sql
- id (UUID, Primary Key)
- name (String, Required)
- email (String, Unique, Required)
- password_hash (String, Required)
- division (String)
- phone (String)
- created_at (Timestamp)
```

### Servers Table
```sql
- id (UUID, Primary Key)
- name (String, Required)
- url (String, Required)
- status (String: UP/DOWN/UNKNOWN)
- response_time (BigInt)
- last_checked (Timestamp)
- created_by (UUID, Foreign Key)
- created_at (Timestamp)
```

### Server Down History Table
```sql
- id (UUID, Primary Key)
- server_id (UUID, Foreign Key)
- server_name (String, Required)
- url (String, Required)
- status (String: DOWN/RESOLVED)
- timestamp (Timestamp)
- created_by (UUID, Foreign Key)
- description (String)
- resolved_by (UUID, Nullable)
- resolved_at (Timestamp, Nullable)
- resolve_note (String)
- assigned_to (UUID, Required) -- User responsible
```

## 🔒 Security Features

- **JWT Authentication**: Stateless token-based auth
- **Password Hashing**: bcrypt dengan cost factor tinggi
- **Input Validation**: Comprehensive request validation
- **CORS Protection**: Configurable cross-origin policies
- **SQL Injection Prevention**: GORM parameterized queries
- **UUID Primary Keys**: Prevent enumeration attacks

## 📈 Performance & Scalability

- **Connection Pooling**: GORM dengan PostgreSQL connection pool
- **Database Indexing**: Optimized queries dengan proper indexing
- **Caching Ready**: Architecture prepared untuk Redis caching
- **Horizontal Scaling**: Stateless design untuk load balancing
- **Monitoring**: Structured logging untuk observability

## 🚀 Deployment Checklist

- [ ] Environment variables configured
- [ ] Firebase service account key placed
- [ ] PostgreSQL database created and migrated
- [ ] SSL certificates untuk production
- [ ] Reverse proxy (nginx) configured
- [ ] Monitoring tools (Prometheus/Grafana) setup
- [ ] Backup strategy implemented
- [ ] CI/CD pipeline configured

## 🗂️ Project Structure

```
NetGuardServer/
├── config/           # Configuration management
├── controllers/      # HTTP request handlers (pure functions)
├── dto/             # Data Transfer Objects (request/response)
├── middleware/      # Custom middleware (JWT auth, CORS)
├── models/          # Database models (GORM)
├── repository/      # Data access layer (interfaces + implementations)
├── routes/          # Route definitions & dependency injection setup
├── services/        # Business logic layer (interfaces + implementations)
├── utils/           # Utility functions (JWT, password, validation, response)
├── di/              # Dependency injection (Google Wire)
├── main.go          # Application entry point
├── .air.toml        # Air hot reload config
├── docker-compose.yml # Docker orchestration
├── Dockerfile       # Multi-stage Docker build
├── .env.example     # Environment variables template
├── .gitignore       # Git ignore rules
├── API_DOCUMENTATION.md # Complete API reference
└── README.md
```

## 🔒 Security Features

- JWT token authentication
- Password hashing dengan bcrypt
- Input validation dan sanitization
- CORS protection
- SQL injection prevention (GORM)

## 📊 Database Schema

### Users Table
```sql
- id (UUID, Primary Key)
- name (String)
- email (String, Unique)
- password_hash (String)
- division (String)
- phone (String)
- created_at (Timestamp)
```

### Servers Table
```sql
- id (UUID, Primary Key)
- name (String)
- url (String)
- status (String: UP/DOWN/UNKNOWN)
- response_time (BigInt)
- last_checked (Timestamp)
- created_by (UUID, Foreign Key)
- created_at (Timestamp)
```

### Server Down History Table
```sql
- id (UUID, Primary Key)
- server_id (UUID, Foreign Key)
- server_name (String)
- url (String)
- status (String)
- timestamp (Timestamp)
- created_by (UUID, Foreign Key)
- description (String)
```

## 🚀 Deployment

### Production Build
```bash
# Build untuk production
go build -ldflags="-s -w" -o netguard .

# Jalankan
./netguard
```

### Environment Variables untuk Production
```env
DB_HOST=your_production_db_host
DB_USER=your_production_db_user
DB_PASSWORD=your_production_db_password
JWT_SECRET=your_secure_jwt_secret
FIREBASE_SERVICE_ACCOUNT_PATH=/app/config/serviceAccount.json
```

## 🤝 Contributing

1. Fork the repository
2. Create feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to branch (`git push origin feature/AmazingFeature`)
5. Open Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Support

For support, email support@uniguard.co.id or create an issue in this repository.

---

**NetGuard Backend** - Real-time server monitoring with mobile push notifications 🚀