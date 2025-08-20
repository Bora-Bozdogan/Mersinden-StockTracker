# Mersinden Women's Cooperatives Product Stock Tracking System

This project is a fullstack **product stock management application** designed for women’s cooperatives in Mersin.  
It provides merchants with a web-based dashboard to manage their products, update stock information, and maintain merchant details.  

The system is built with:
- **Frontend**: HTML, CSS, and JavaScript  
- **Backend**: Go (Fiber web framework) with PostgreSQL  
- **Authentication**: Firebase Authentication  
- **ORM**: GORM  

---

## 🚀 Features

- **User Authentication** (via Firebase)  
- **Merchant Management**  
  - View merchant info  
  - Update merchant details  
- **Product Management**  
  - Add new products (name, description, price, stock)  
  - Update product details  
  - Delete products  
  - Role-based access: admins can see all items, merchants only see their own  
- **Modern UI**  
  - Responsive login and dashboard pages  
  - Styled with CSS (custom theme + Google Fonts)  

---

## 📂 Project Structure
.
├── frontend/
│   ├── dashboard.html       # Dashboard UI
│   ├── dashboard.css        # Dashboard styles
│   ├── dashboard.js         # Dashboard functionality
│   ├── login.css            # Login page styles
│   ├── login.js             # Login functionality
│   └── index.js             # Entry point script (frontend)
│
├── backend/
│   ├── main.go              # Main Fiber server setup
│   ├── handlers.go          # HTTP request handlers
│   ├── services.go          # Business logic layer
│   ├── product_repository.go# Repository for products
│   ├── ... (merchant repo, config, firebase client, models)
│
└── README.md
---

## 🛠️ Setup & Installation

### 1. Prerequisites
- Go 1.20+  
- Node.js (if extending frontend build)  
- PostgreSQL  
- Firebase project (for authentication)  

### 2. Backend Setup
1. Clone the repository
   git clone https://github.com/your-org/mersinden-stockapp.git
   cd mersinden-stockapp/backend

2. Configure environment variables in a config file:
   - PostgreSQL connection details  
   - Firebase credentials file path  
   - Server listen port and frontend address  

3. Run the backend:
   go run main.go

### 3. Frontend Setup
Open `dashboard.html` or `login.html` in a browser (or serve them with a simple HTTP server).  
They connect to the backend API routes.

---

## 🔗 API Endpoints

### Items
- GET /items → List items for the authenticated merchant (or all if admin)  
- POST /items → Create new product  
- PUT /items/:id → Update product  
- DELETE /items/:id → Delete product  

### Merchant
- GET /merchant/:id → Get merchant by ID  
- GET /merchant → Get current merchant (by UID)  
- PUT /merchant → Update merchant info  

All routes require Firebase-authenticated requests.

---

## 🎨 UI Overview

- **Login Page** (login.css)  
  Clean, card-based login with error handling.  
- **Dashboard Page** (dashboard.html, dashboard.css)  
  - Product listing with add/update/delete  
  - Merchant data panel with editable fields  
  - Floating logout button  

---

## 👩‍💻 Tech Stack

- **Frontend**: Vanilla JS, HTML, CSS  
- **Backend**: Go (Fiber, GORM)  
- **Database**: PostgreSQL  
- **Auth**: Firebase Authentication  

---

## 📜 License

This project is **not free to use or modify**.  
It is shared **only for observation and reference purposes**.
