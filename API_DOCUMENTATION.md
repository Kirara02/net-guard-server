# NetGuard API Documentation

## 📡 Complete API Reference

This document provides detailed information about all NetGuard Backend API endpoints, including request/response formats, authentication requirements, and usage examples.

## 🔐 Authentication Endpoints

### **POST /api/auth/register**

Register new user account

**Request Body:**

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123",
  "division": "IT",
  "phone": "08123456789",
  "role": "USER"
}
```

**Response (201):**

```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": "uuid",
      "name": "John Doe",
      "email": "john@example.com",
      "division": "IT",
      "phone": "08123456789",
      "role": "USER",
      "is_active": true,
      "created_at": "2024-01-01T00:00:00Z"
    }
  }
}
```

### **POST /api/auth/login**

User authentication

**Request Body:**

```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response (200):**

```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": "uuid",
      "name": "John Doe",
      "email": "john@example.com",
      "role": "USER"
    }
  }
}
```

### **GET /api/auth/me**

Get current user profile

**Headers:**

```
Authorization: Bearer <jwt_token>
```

**Response (200):**

```json
{
  "success": true,
  "data": {
    "id": "uuid",
    "name": "John Doe",
    "email": "john@example.com",
    "division": "IT",
    "phone": "08123456789",
    "role": "USER",
    "is_active": true,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

## 🌐 Server Management Endpoints

### **POST /api/servers**

Create new server

**Headers:**

```
Authorization: Bearer <jwt_token>
```

**Request Body:**

```json
{
  "name": "API Server",
  "url": "https://api.company.com"
}
```

**Response (200):**

```json
{
  "success": true,
  "message": "Server created successfully",
  "data": {
    "id": "uuid",
    "name": "API Server",
    "url": "https://api.company.com",
    "created_by": "user-uuid",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

### **GET /api/servers**

Get all servers from all users

**Headers:**

```
Authorization: Bearer <jwt_token>
```

**Response (200):**

```json
{
  "success": true,
  "data": [
    {
      "id": "uuid",
      "name": "API Server",
      "url": "https://api.company.com",
      "created_by": "user-uuid",
      "created_at": "2024-01-01T00:00:00Z"
    },
    {
      "id": "uuid-2",
      "name": "Database Server",
      "url": "https://db.company.com",
      "created_by": "user-uuid-2",
      "created_at": "2024-01-02T00:00:00Z"
    }
  ]
}
```

### **GET /api/servers/:id**

Get specific server by ID

**Headers:**

```
Authorization: Bearer <jwt_token>
```

**Response (200):**

```json
{
  "success": true,
  "data": {
    "id": "uuid",
    "name": "API Server",
    "url": "https://api.company.com",
    "created_by": "user-uuid",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

### **PUT /api/servers/:id**

Update server information

**Headers:**

```
Authorization: Bearer <jwt_token>
```

**Request Body:**

```json
{
  "name": "Updated API Server",
  "url": "https://api-v2.company.com"
}
```

**Response (200):**

```json
{
  "success": true,
  "message": "Server updated successfully",
  "data": {
    "id": "uuid",
    "name": "Updated API Server",
    "url": "https://api-v2.company.com",
    "created_by": "user-uuid",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

### **DELETE /api/servers/:id**

Delete server

**Headers:**

```
Authorization: Bearer <jwt_token>
```

**Response (200):**

```json
{
  "success": true,
  "message": "Server deleted successfully",
  "data": null
}
```

### **PATCH /api/servers/:id/status**

Update server status (Mobile App Endpoint)

**Headers:**

```
Authorization: Bearer <jwt_token>
```

**Request Body:**

```json
{
  "status": "DOWN",
  "response_time": 5000
}
```

**Response (200):**

```json
{
  "success": true,
  "message": "Server status updated successfully",
  "data": {
    "id": "uuid",
    "name": "API Server",
    "url": "https://api.company.com",
    "created_by": "user-uuid",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

## 🚨 Incident Management (History) Endpoints

### **POST /api/history**

Create incident record (Auto-triggered by status update)

**Headers:**

```
Authorization: Bearer <jwt_token>
```

**Request Body:**

```json
{
  "server_id": "server-uuid",
  "server_name": "API Server",
  "url": "https://api.company.com",
  "status": "DOWN"
}
```

**Response (200):**

```json
{
  "success": true,
  "message": "History record created successfully",
  "data": {
    "id": "uuid",
    "server_id": "server-uuid",
    "server_name": "API Server",
    "url": "https://api.company.com",
    "status": "DOWN",
    "timestamp": "2024-01-01T00:00:00Z",
    "created_by": "user-uuid"
  }
}
```

### **GET /api/history**

Get all incident history from all users

**Headers:**

```
Authorization: Bearer <jwt_token>
```

**Query Parameters:**

- `server_id` (optional): Filter by server
- `limit` (optional): Limit results (default: 50, max: 1000)

**Response (200):**

```json
{
  "success": true,
  "data": [
    {
      "id": "uuid",
      "server_id": "server-uuid",
      "server_name": "API Server",
      "url": "https://api.company.com",
      "status": "DOWN",
      "timestamp": "2024-01-01T00:00:00Z",
      "created_by": "John Doe",
      "resolved_by": null,
      "resolved_at": null,
      "resolve_note": null,
      "description": null
    },
    {
      "id": "uuid-2",
      "server_id": "server-uuid-2",
      "server_name": "Database Server",
      "url": "https://db.company.com",
      "status": "RESOLVED",
      "timestamp": "2024-01-02T00:00:00Z",
      "created_by": "Bob Wilson",
      "resolved_by": "Alice Johnson",
      "resolved_at": "2024-01-02T01:30:00Z",
      "resolve_note": "Database connection restored",
      "description": null
    }
  ]
}
```

### **PATCH /api/history/:id/resolve**

Resolve incident

**Headers:**

```
Authorization: Bearer <jwt_token>
```

**Request Body:**

```json
{
  "resolve_note": "Server restarted, issue resolved"
}
```

**Response (200):**

```json
{
  "success": true,
  "message": "History record resolved successfully",
  "data": {
    "id": "uuid",
    "server_id": "server-uuid",
    "server_name": "API Server",
    "url": "https://api.company.com",
    "status": "RESOLVED",
    "timestamp": "2024-01-01T00:00:00Z",
    "created_by": "user-name",
    "resolved_by": "user-name",
    "resolved_at": "2024-01-01T01:00:00Z",
    "resolve_note": "Server restarted, issue resolved"
  }
}
```

### **GET /api/history/report/monthly**

Get monthly server down report

**Headers:**

```
Authorization: Bearer <jwt_token>
```

**Query Parameters:**

- `year` (required): Year (e.g., 2024)
- `month` (required): Month (1-12)

**Response (200):**

```json
{
  "success": true,
  "message": "Monthly report generated successfully",
  "data": {
    "year": 2024,
    "month": 10,
    "report": [
      {
        "server_id": "server-uuid",
        "server_name": "API Server",
        "url": "https://api.company.com",
        "down_count": 5,
        "resolved_count": 4,
        "avg_resolution_time": 3600.5
      }
    ]
  }
}
```

## 📊 Error Response Format

**All error responses follow this format:**

```json
{
  "success": false,
  "error": "Error message description"
}
```

**Common HTTP Status Codes:**

- `200` - Success
- `201` - Created
- `400` - Bad Request (validation error)
- `401` - Unauthorized (invalid/missing token)
- `403` - Forbidden (insufficient permissions)
- `404` - Not Found
- `409` - Conflict (duplicate data)
- `500` - Internal Server Error

## 🔑 Authentication

All protected endpoints require JWT token in Authorization header:

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

## 📱 Mobile App Integration

### Status Update Flow:

1. Mobile app detects server DOWN
2. Call `PATCH /api/servers/:id/status` with status "DOWN"
3. Backend automatically creates history record
4. FCM notification sent to all users
5. Users can resolve incidents via mobile app

### FCM Notification Payload:

```json
{
  "notification": {
    "title": "Server DOWN: API Server",
    "body": "https://api.company.com"
  },
  "data": {
    "server_id": "uuid",
    "status": "DOWN"
  }
}
```

## 🧪 Testing Examples

### Register User

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123"
  }'
```

### Create Server

```bash
curl -X POST http://localhost:8080/api/servers \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "API Server",
    "url": "https://api.company.com"
  }'
```

### Update Server Status (Mobile App)

```bash
curl -X PATCH http://localhost:8080/api/servers/SERVER_ID/status \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "status": "DOWN",
    "response_time": 5000
  }'
```

### Resolve Incident

```bash
curl -X PATCH http://localhost:8080/api/history/HISTORY_ID/resolve \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "resolve_note": "Server restarted successfully"
  }'
```

### Get Monthly Report

```bash
curl -X GET "http://localhost:8080/api/history/report/monthly?year=2024&month=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

**Last Updated:** October 18, 2024
**Version:** 1.0.0
