🐾 Pet Clinic Management System (Go + PostgreSQL)

A complete backend project built using Golang, PostgreSQL, JWT Authentication, Role-Based Access(Staff / Owner), Structured Logging, File Upload & Download, and Postman Collection.

This project fulfills the full requirements of your To-Do Pet Clinic Assignment.

🚀 Features
✅ Owners Module

Create Owner

Get All Owners

Update Owner

Delete Owner

Validation (name, email required)

Logging for all operations

✅ Pets Module

Add Pet

Get Pets

Update Pet

Delete Pet

Owners can only manage their own pets

Staff can manage all pets

✅ Appointments

Book Appointment

View Appointments

Update Appointment

Cancel Appointment

✅ Authentication (JWT)

Two roles:

Staff → Full access

Owner → Limited to own pets

Login using:

Basic Auth

OR JSON Body

JWT Middleware protects all /api/... routes

✅ Logging + Error Handling

Using Logrus + Lumberjack:

INFO logs

DEBUG logs

WARN logs

ERROR logs

Log rotation

All errors return clean JSON responses

✅ File Management

Upload medical records (PDF / images)

Download files securely

Dedicated /uploads folder

MIME-type handled

Logging for upload/download events


🗂 Database Design (PostgreSQL)

owners

id SERIAL PRIMARY KEY
name VARCHAR
contact VARCHAR
email VARCHAR UNIQUE


pets

id SERIAL PRIMARY KEY
name VARCHAR
species VARCHAR
breed VARCHAR
owner_id INT REFERENCES owners(id)
medical_history TEXT


appointments

id SERIAL PRIMARY KEY
date DATE
time TIME
pet_id INT REFERENCES pets(id)
reason TEXT


Sample data included in database.sql.


🛠 Tech Stack

Go 1.21+

PostgreSQL

gorilla/mux

JWT (golang-jwt v5)

Logrus

Lumberjack (log rotation)

Postman (API testing)


📦 Project Structure
pet-clinic/
│── auth/               (JWT + middleware)
│── handlers/           (All API handlers)
│── models/             (Struct models)
│── utils/              (Logger setup)
│── uploads/            (Uploaded files)
│── main.go
│── go.mod
│── database.sql
│── postman_collection.json
│── README.md
│── .gitignore


🔧 Setup Instructions
1️⃣ Clone the Repo
git clone https://github.com/Shreyas071845/pet-clinic.git
cd pet-clinic

2️⃣ Install Dependencies
go mod tidy

3️⃣ Create .env File

Create a file named .env:

JWT_SECRET=mysecretkey

4️⃣ Import Database

Run PostgreSQL command:

psql -U postgres -d petclinic -f database.sql

5️⃣ Run the Server
go run main.go


Server runs at:

http://localhost:8080

📬 Postman Collection

A ready-to-use Postman collection is included:

postman_collection.json


Import it in Postman and start testing immediately.

🔐 How Authentication Works
Staff Login

POST → /login

{
  "username": "staff1",
  "password": "staffpass"
}

Owner Login

POST → /login

{
  "username": "owner1",
  "password": "ownerpass"
}


The response will give a JWT token.

Use Token for All Protected Routes

In Postman → Authorization:

Bearer Token
{{token}}


Token is auto-saved via Post-response script.

📤 File Upload

Endpoint:

POST /api/upload


Body → form-data:

file: <select your PDF/image>

📥 File Download
GET /api/download/<filename>


Example:

GET http://localhost:8080/api/download/test.pdf

🧑‍💻 Author

Shreyas Bhat
GitHub: Shreyas071845

🎉 Final Notes

This project is complete with:
✔ Authentication
✔ CRUD operations
✔ Logging
✔ File management
✔ Role-based access
✔ Postman support
✔ Clean structured Go code