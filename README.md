#🐾 Pet Clinic Management System – Go + PostgreSQL

A backend project built using Golang, PostgreSQL, JWT authentication, Role-Based Access (Staff/Owner), Logging, and File Upload/Download.
This project was developed as part of a training assignment to understand API creation, authentication, database handling, and file management.

##📜 Short Description / Purpose

The Pet Clinic Management System allows pet owners and staff to manage owners, pets, and appointments.
It includes secure login, validation, structured logging, and database storage.
Owners can manage only their pets, and staff have full access to all data.
---
##🛠️ Tech Stack

This system uses the following tools and technologies:

- Golang (net/http, gorilla/mux) – Backend API development

- PostgreSQL – Database for storing owners, pets, and appointments

- JWT (HS256) – Authentication and authorization

- File Upload/Download – For medical reports (PDFs/images)

- Logrus + Lumberjack – Structured logging with log rotation

- Postman – API testing

---

- SQL – Table creation and sample data

##🌐 Data Flow / Modules

- Owners – Add, view, update, delete

- Pets – Add, view, update, delete

- Appointments – Book, view, update, cancel

- Authentication – JSON or Basic login → JWT token

**Role-Based Access**

- Owner → only own pets

- Staff → all pets

- File Management

- Upload medical files

- Download stored files
---
✨ Features / Highlights

JWT-based secure login

Role-based authorization (Owner/Staff)

Structured logging (INFO, DEBUG, WARN, ERROR)

CRUD operations for owners, pets, appointments

Input validation for important fields

File upload & download support

Database integration with PostgreSQL

Postman collection included

🧩 Business Use Cases / Purpose

This system can be used for:

Clinics – Managing pet records, appointments, and medical history

Owners – Tracking their pets’ visits and medical details

Training purposes – Understanding backend systems, JWT, databases, and file handling

📂 Repository Contents

auth/ – JWT generation and middleware

handlers/ – API endpoints for owners, pets, appointments, files

db/ – PostgreSQL connection

models/ – Data models

utils/ – Logger setup

database.sql – SQL tables + sample data

postman_collection.json – Ready-to-use Postman import

uploads/ – File storage (ignored in Git)

main.go – Main server entry

README.md – Documentation

🚀 Getting Started

Clone the repository

git clone https://github.com/Shreyas071845/pet-clinic.git
cd pet-clinic


Install dependencies

go mod tidy


Create .env file

JWT_SECRET=mysecretkey


Import database

psql -U postgres -d petclinic -f database.sql


Run server

go run main.go


Server will be available at:
http://localhost:8080

🔐 Authentication
Staff Login
POST /login
{
  "username": "staff1",
  "password": "staffpass"
}

Owner Login
POST /login
{
  "username": "owner1",
  "password": "ownerpass"
}


Use the returned JWT as:
Authorization: Bearer <token>

📤 File Upload
POST /api/upload


Body → form-data → file: <choose file>

📥 File Download
GET /api/download/<filename>

🖼️ Preview (Optional Screenshot Section)

If you want, we can add sample screenshots of Postman, database, or folder structure here.

🧑‍💻 Author

Shreyas Bhat
GitHub: Shreyas071845
