# Ecommerce Project (Go)

A comprehensive ecommerce backend built with Go. This project provides RESTful APIs for managing products, users, orders, and payments.

## Features

- User authentication & authorization (JWT)
- Product catalog management (CRUD)
- Shopping cart functionality
- Order processing
- Payment integration (mocked or real)
- Admin dashboard endpoints
- Secure password hashing
- Pagination & filtering
- Error handling & validation

## Technologies

- Go (Golang)
- Gin (HTTP web framework)
- GORM (ORM for database)
- PostgreSQL/MySQL/SQLite (configurable)
- JWT (authentication)
- Docker (optional)

## Getting Started

### Prerequisites

- Go 1.20+
- PostgreSQL/MySQL/SQLite
- Git

### Installation

```bash
git clone https://github.com/yourusername/ecommerce-project.git
cd ecommerce-project
go mod tidy
```

### Configuration

Edit `config.yaml` or set environment variables for database and JWT secret.

### Running

```bash
go run main.go
```

### Docker

```bash
docker-compose up --build
```

## API Endpoints

| Method | Endpoint             | Description                |
|--------|----------------------|----------------------------|
| POST   | /api/register        | Register new user          |
| POST   | /api/login           | User login                 |
| GET    | /api/products        | List products              |
| POST   | /api/products        | Add product (admin)        |
| PUT    | /api/products/:id    | Update product (admin)     |
| DELETE | /api/products/:id    | Delete product (admin)     |
| POST   | /api/cart            | Add to cart                |
| GET    | /api/cart            | View cart                  |
| POST   | /api/orders          | Place order                |
| GET    | /api/orders          | List user orders           |

## Folder Structure

```
├── cmd/
├── config/
├── controllers/
├── models/
├── routes/
├── middleware/
├── utils/
├── main.go
└── README.md
```

## Contributing

1. Fork the repo
2. Create your feature branch (`git checkout -b feature/foo`)
3. Commit your changes
4. Push to the branch
5. Open a pull request

## License

MIT

## Contact

For questions, open an issue or email [your.email@example.com](mailto:your.email@example.com).