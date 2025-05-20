1- Architecture overview: 

[ Web UI (HTML + Bootstrap) ]
         ↓
[ Gin HTTP Server ]
         ↓
[ Auth | Message | Media | Swagger Docs ]
         ↓
[ GORM ORM + PostgreSQL DB ]
         ↓
[ Dockerized Deployment ]


2- Tech stack:

Lang : GO
DB: Postgree

3- Setup Instructions

  1. Clone the repository git clone https://github.com/emaanmohamed/chat-app.git
  2. Copy env example into your env
  3. docker-compose up --build

4- Access the app
  1- login page http://localhost:8083/static/index.html
  2- signup page http://localhost:8083/static/signup.html
  3- chat UI http://localhost:8083/static/chat.html
