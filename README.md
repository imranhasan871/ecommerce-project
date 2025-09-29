# E-commerce Product API

A lightweight e-commerce REST API built with Go that provides product management functionality with CORS support and request logging.

## 🏗️ Architecture

This project follows a clean, modular architecture with the following structure:

```
ecommerce-project/
├── global_router/          # Global routing and CORS handling
├── handlers/              # HTTP request handlers
├── middleware/            # HTTP middleware (logging)
├── product/              # Product domain models
├── utils/                # Utility functions
├── main.go               # Application entry point
├── docker-compose.yaml   # Multi-service Docker setup
└── Dockerfile           # Go application containerization
```

## 🚀 Features

- **Product Management**: Create and retrieve products
- **RESTful API**: Clean REST endpoints with proper HTTP methods
- **CORS Support**: Cross-origin requests enabled for web applications
- **Request Logging**: Automatic logging of all HTTP requests with timing
- **Dockerized**: Ready-to-deploy with Docker and Docker Compose
- **Database Ready**: PostgreSQL setup included in docker-compose

## 📋 API Endpoints

### Products

| Method | Endpoint    | Description              | Body Required |
|--------|-------------|--------------------------|---------------|
| GET    | `/products` | Get all products         | No            |
| POST   | `/products` | Create a new product     | Yes           |

### Product Model

```json
{
  "id": 1,
  "title": "Product Name",
  "description": "Product description",
  "price": 99.99,
  "imageURL": "https://example.com/image.jpg"
}
```

## 🛠️ Installation & Setup

### Prerequisites

- Go 1.25+
- Docker & Docker Compose (optional)

### Local Development

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd ecommerce-project
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run the application**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:3000`

### Docker Setup

1. **Start all services**
   ```bash
   docker-compose up -d
   ```

This will start:
- **Go Application**: `http://localhost:8081`
- **PostgreSQL Database**: `localhost:5432`
- **pgAdmin**: `http://localhost:8080`

2. **Database credentials**
   - **User**: admin
   - **Password**: adminpass
   - **Database**: ecommerce_db

3. **pgAdmin credentials**
   - **Email**: admin@admin.com
   - **Password**: adminpass

## 📡 Usage Examples

### Get All Products

```bash
curl -X GET http://localhost:3000/products
```

**Response:**
```json
[
  {
    "id": 1,
    "title": "Wireless Headphones",
    "description": "High-quality noise-cancelling headphones.",
    "price": 129.99,
    "imageURL": "https://example.com/headphones.jpg"
  },
  {
    "id": 2,
    "title": "Smart Watch",
    "description": "Stylish smart watch with health tracking.",
    "price": 199.99,
    "imageURL": "https://example.com/watch.jpg"
  }
]
```

### Create a Product

```bash
curl -X POST http://localhost:3000/products \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Gaming Mouse",
    "description": "High-precision gaming mouse",
    "price": 79.99,
    "imageURL": "https://example.com/mouse.jpg"
  }'
```

**Response:**
```json
{
  "id": 7,
  "title": "Gaming Mouse",
  "description": "High-precision gaming mouse",
  "price": 79.99,
  "imageURL": "https://example.com/mouse.jpg"
}
```

## 🏛️ Code Structure

### Core Components

#### `main.go`
- Application entry point
- HTTP server setup on port 3000
- Route registration
- Sample data initialization

#### `product/product.go`
- Product struct definition with JSON tags
- Core product model

#### `product/product_list.go`
- In-memory product storage
- Global product slice

#### `handlers/`
- **`get_products.go`**: Returns all products from the product list
- **`create_product.go`**: Creates new products with auto-generated IDs

#### `middleware/logger.go`
- HTTP request logging middleware
- Tracks request method, path, and response time

#### `global_router/global_router.go`
- CORS headers configuration
- OPTIONS request handling
- Global request wrapper

#### `utils/send_data.go`
- JSON response utility function
- Centralized data serialization

## 🔧 Configuration

### Environment Variables (Docker)

| Variable     | Description          | Default Value |
|-------------|----------------------|---------------|
| `DB_HOST`   | Database host        | postgres      |
| `DB_PORT`   | Database port        | 5432          |
| `DB_USER`   | Database username    | admin         |
| `DB_PASSWORD` | Database password  | adminpass     |
| `DB_NAME`   | Database name        | ecommerce_db  |

### CORS Configuration

The application is configured to allow:
- **Origins**: All (`*`)
- **Methods**: GET, POST, PUT, DELETE, OPTIONS
- **Headers**: Content-Type

## 🚧 Current Implementation

- **Storage**: In-memory product list (initialized with 6 sample products)
- **Database**: PostgreSQL configured but not yet integrated
- **Authentication**: Not implemented
- **Validation**: Basic JSON validation

## 🔮 Future Enhancements

- [ ] Database integration with PostgreSQL
- [ ] User authentication and authorization
- [ ] Product categories and filtering
- [ ] Inventory management
- [ ] Order processing
- [ ] Payment integration
- [ ] Advanced search functionality
- [ ] Image upload handling
- [ ] API documentation with Swagger

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

---

**Built with ❤️ using Go**