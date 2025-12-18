# Go Lang CRUD (File-Based)

This project is a **simple CRUD (Create, Read, Update, Delete) REST API** built using **Go (Golang)** and **Gorilla Mux**.

It is designed for **learning purposes** and demonstrates how CRUD works **without using a database**, by storing data in a simple **text file (`users.txt`)**.

---

## 🚀 Features

- REST API using Go
- CRUD operations (Create, Read, Update, Delete)
- File-based storage (`users.txt`)
- Auto-increment user ID
- Beginner-friendly project structure
- No database required

---

## 🧱 Project Structure

go-lang-CRUD/
│── main.go # Server & routes
│── handlers.go # CRUD handlers
│── models.go # User model
│── file_store.go # File read/write logic
│── users.txt # Data storage (gitignored)
│── go.mod
│── go.sum
│── .gitignore
│── README.md

yaml
Copy code

---

## ▶️ How to Run the Project

### 1️⃣ Clone the repository
```bash
git clone https://github.com/ankurchouhan/go-lang-CRUD.git
cd go-lang-CRUD
2️⃣ Run the server
bash
Copy code
go run .
Server will start at:

arduino
Copy code
http://localhost:8080
📌 CRUD API Usage (Paste in Terminal)
Make sure the server is running before executing commands.

➕ Create User
powershell
Copy code
Invoke-RestMethod `
 -Uri http://localhost:8080/users `
 -Method POST `
 -ContentType "application/json" `
 -Body '{"name":"Ankur","email":"ankur@mail.com"}'
📄 Get All Users (READ)
powershell
Copy code
Invoke-RestMethod http://localhost:8080/users
🔍 Get User By ID (READ)
powershell
Copy code
Invoke-RestMethod http://localhost:8080/users/1
✏️ Update User (EDIT)
powershell
Copy code
Invoke-RestMethod `
 -Uri http://localhost:8080/users/1 `
 -Method PUT `
 -ContentType "application/json" `
 -Body '{"name":"Ankur Kumar","email":"ankur.kumar@mail.com"}'
✔ Updates user with ID 1
✔ Changes are saved to users.txt

🗑️ Delete User
powershell
Copy code
Invoke-RestMethod `
 -Uri http://localhost:8080/users/1 `
 -Method DELETE
✔ Deletes the user permanently from file storage

📂 Data Storage
All user data is stored in a text file:

Copy code
users.txt
File format:

bash
Copy code
id,name,email
Example:

graphql
Copy code
1,Ankur,ankur@mail.com
2,Rahul,rahul@mail.com
⚠️ users.txt is ignored using .gitignore and is not committed to GitHub.
