# GH Gist Blog

A modern blog application built with Go, featuring JWT authentication, role-based access control, and MongoDB integration and the Gin framework.This project is to help me learn how blog application backends are structured and built.

## Status

Work in Progress - Core authentication and user management implemented

## Features

### Implemented
- JWT-based authentication
- Role-based access control (Publisher, Writer, Reader)
- User registration and login
- Password hashing with bcrypt
- Input validation (email, password, username)
- MongoDB integration with connection pooling
- Docker containerization
- Environment-based configuration

### Planned
- [ ] Article creation and management
- [ ] Category system
- [ ] File upload integration
- [ ] Search functionality
- [ ] User profile management
- [ ] API documentation

## Tech Stack

- **Backend**: Go (Gin framework)
- **Database**: MongoDB
- **Authentication**: JWT tokens
- **Containerization**: Docker & Docker Compose
- **Password Hashing**: bcrypt

## Project Structure

```
gh gist blog/
├── Database/           # Database connection
├── handlers/           # HTTP handlers and routes
├── middleware/         # Authentication middleware
├── models/            # Data models
├── repository/        # Database operations
├── services/          # Business logic
├── utils/             # Utility functions (JWT, password)
├── validation/        # Input validation
├── docker-compose.yml # Docker configuration
├── Dockerfile         # Container build
└── main.go           # Application entry point
```

## Getting Started

### Prerequisites

- Go 1.19 or higher
- MongoDB (local or Atlas)
- Docker (optional)

### Environment Variables

Create a `.env` file in the root directory:

```env
SECRET_KEY=your-jwt-secret-key
PORT=8080
APP_URL=http://localhost:8080
MONGODB_URI=mongodb://localhost:27018/ghgistDB
```

### Local Development

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Start MongoDB (if running locally)
4. Run the application:
   ```bash
   go run main.go
   ```

### Docker Development

1. Start all services:
   ```bash
   docker compose up --build
   ```

## API Endpoints

### Public Routes
- `GET /api/public/writers` - Fetch all writers
- `POST /api/public/auth/login` - User login
- `POST /api/public/auth/register` - User registration (temporary public access)

### Protected Routes
- `POST /api/admin/articles` - Create article (requires authentication)

### Role-Based Routes
- `POST /api/admin/auth/register` - Create user (publisher only)

## Authentication

The application uses JWT tokens for authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

### User Roles

- **Publisher**: Can create users and articles
- **Writer**: Can create articles
- **Reader**: Read-only access

## Development Workflow

- **Local Development**: Use `go run main.go` for fast iteration
- **Container Testing**: Use `docker compose up --build` before commits
- **Production**: Deploy using Docker containers

## Deployment

### Fly.io
- Uses Dockerfile for deployment
- Set environment variables via `fly secrets`
- MongoDB Atlas recommended for production database

### Environment-Specific Configuration
- **Local**: Uses `.env` file
- **Docker**: Uses docker-compose environment variables
- **Production**: Uses platform-specific environment variables

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test with Docker before committing
5. Submit a pull request

## License

[Add your license here]

## Contact

[Add your contact information here]