Go + Gin + MongoDB (User CRUD)

A simple backend project to understand how MongoDB integrates with Go using Gin.
This project focuses on clean structure, basic CRUD operations, and good coding practices.

🚀 Tech Stack
Go (Golang)
Gin (HTTP framework)
MongoDB
Air (hot reload for development)
📁 Project Structure
go-mongo/
├── cmd/                # Entry point (main.go)
├── internal/
│   ├── controller/    # HTTP handlers
│   ├── service/       # Business logic
│   ├── repository/    # DB interactions
│   └── db/            # MongoDB connection
├── model/             # Data models
├── go.mod
⚙️ Features
Create User
Get User by ID
Get All Users
Update User
Clean layered architecture (Handler → Service → Repository)
MongoDB integration with proper ObjectID handling
🔧 Setup & Run
1. Clone the repo
git clone <your-repo-link>
cd go-mongo
2. Set environment variable

Create .env or set manually:

MONGO_URI=your_mongodb_connection_string
3. Install dependencies
go mod tidy
4. Run project

Using Air (recommended):

air

Or normally:

go run ./cmd
📌 API Endpoints
Method	Route	Description
GET	/	Health check
POST	/newuser	Create user
GET	/user/:id	Get user by ID
GET	/users	Get all users
PATCH	/user/:id	Update user
🧠 Learnings
How MongoDB works with Go (BSON, ObjectID, queries)
Structuring backend apps cleanly
Handling partial updates safely
Using Gin for fast API development
Using Air for hot reload (better dev experience)
📌 Note

This is a learning project focused on understanding integration and structure.
It will be used as a base for integrating MongoDB into a larger AI task processor project.

🚀 Future Improvements
Add authentication
Add validation layer
Pagination & filtering
Docker setup
👨‍💻 Author

Vivek