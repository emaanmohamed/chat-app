# 1. Architecture Overview

[ Web UI (HTML + Bootstrap) ]
↓
[ Gin HTTP Server ]
↓
[ Auth | Message | Media | Swagger Docs ]
↓
[ GORM ORM + PostgreSQL DB ]
↓
[ Dockerized Deployment ]

# 2. Tech Stack

- **Language:** Go
- **Database:** PostgreSQL

---

# 3. Setup Instructions

1. Clone the repository:
   [https://github.com/emaanmohamed/chat-app.git](https://github.com/emaanmohamed/chat-app.git)
2. Copy `.env.example` to `.env`
3. Run:
   ```bash
   docker-compose up --build

# 4- Access the app

Login Page: http://localhost:8083/index.html

Signup Page: http://localhost:8083/signup.html

Chat UI: http://localhost:8083/chat.html