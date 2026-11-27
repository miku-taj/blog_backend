# Blog System API

A RESTful backend API for a multi-user blogging platform built with **Go**, **Gin**, and **PostgreSQL**. Supports user authentication, blog post management, nested or post-level comments, likes on posts and comments, user profiles, and creator subscriptions.

## Tech Stack

* **Language:** Go 1.22.5
* **Web Framework:** [Gin](https://github.com/gin-gonic/gin) (`v1.10.0`)
* **ORM & Database:** [GORM](https://gorm.io/) (`v1.25.12`) with PostgreSQL driver (`pgx/v5`)
* **Authentication & Security:** JWT (`golang-jwt/jwt/v5`) and `bcrypt` password hashing (`golang.org/x/crypto`)

## Features

* **Authentication & Security:** User registration and login with hashed passwords, JWT token verification middleware, CORS configuration, and panic recovery middleware.
* **Posts & Profiles:** Create, read, update, and delete blog posts; manage user profiles.
* **Social Interactions:** Comment on posts, like/unlike both posts and comments, and subscribe/unsubscribe to other users.
* **Clean Layered Architecture:** Separated into HTTP handlers, business logic services, GORM repositories, and domain models.

## Project Structure

```text
blog-system/
├── cmd/
│   └── app/
│       └── main.go                 # Application entry point & server initialization
├── internal/
│   ├── handler/                    # HTTP handlers (Gin routes & request binding)
│   │   ├── middleware/
│   │   │   ├── CORS.go             # Cross-Origin Resource Sharing policy
│   │   │   ├── authentication.go   # JWT token validation middleware
│   │   │   └── recovery.go         # Panic recovery middleware
│   │   ├── comment.go
│   │   ├── handler.go
│   │   ├── like.go
│   │   ├── post.go
│   │   ├── profile.go
│   │   ├── subscribe.go
│   │   └── user.go
│   ├── model/                      # GORM database models & entities
│   │   ├── comment.go
│   │   ├── comment_like.go
│   │   ├── post.go
│   │   ├── post_like.go
│   │   ├── subscription.go
│   │   └── user.go
│   ├── repository/                 # Data access layer (PostgreSQL queries)
│   ├── service/                    # Core business logic layer
│   └── utils/
│       ├── hash_password.go        # Password hashing & verification
│       └── jwt_token.go            # JWT generation & parsing
├── sql/
│   └── query.sql                   # SQL schema / queries
├── go.mod
└── go.sum

