# Enterprise Management System - Server Application

## Project Overview

This project is a coursework assignment aimed at developing a server application using object-oriented programming paradigms. The application allows users to manage the distribution of technical equipment across enterprise workspaces and collect data from measuring devices connected to the service.

The system supports automation and digitization of technological processes and enterprise operations as a whole.

## Technology Stack

- **Language**: Go 1.24.0+ ([https://go.dev/dl/](https://go.dev/dl/))
- **IDE**: Visual Studio Code ([https://code.visualstudio.com/](https://code.visualstudio.com/))
- **Database**: PostgreSQL 12+ ([https://www.postgresql.org/](https://www.postgresql.org/))
- **API Testing**: Postman ([https://www.postman.com/](https://www.postman.com/))
- **Project Template**: [BohdanBoriak/boilerplate-go-back](https://github.com/BohdanBoriak/boilerplate-go-back)

## Project Requirements

The application is divided into implementation stages:

### Stage 1: Project Setup & Authorization ✅ COMPLETED
- Installation of software and project deployment on local environment
- Testing of authorization endpoints (Register, Login, Logout)
- Code familiarization and analysis

**Status**: ✅ Fully implemented

### Stage 2: Organization CRUD ✅ COMPLETED
Implement CRUD functionality for the "Organization" entity representing enterprises.

**Database Fields**:
- `id`: Unique identifier (BIGSERIAL PRIMARY KEY)
- `userId`: Owner of the organization (BIGINT FOREIGN KEY)
- `name`: Organization name (VARCHAR 255)
- `description`: Organization description (TEXT)
- `city`: City location (VARCHAR 255)
- `address`: Physical address (VARCHAR 255)
- `lat`: Latitude coordinate (FLOAT8)
- `lon`: Longitude coordinate (FLOAT8)
- `created_date`: Creation timestamp (TIMESTAMP)
- `updated_date`: Last update timestamp (TIMESTAMP)
- `deleted_date`: Soft delete timestamp (TIMESTAMP, NULL)

**API Endpoints**:
- `POST /api/v1/organizations` - Create new organization
- `GET /api/v1/organizations` - List all organizations (with pagination)
- `GET /api/v1/organizations/{id}` - Get organization by ID
- `PUT /api/v1/organizations/{id}` - Update organization
- `DELETE /api/v1/organizations/{id}` - Delete organization (soft delete)

**Status**: ✅ Fully implemented

### Stage 3: Room CRUD - TO BE IMPLEMENTED
Implement CRUD functionality for the "Room" entity representing workspaces within organizations.

**Database Fields**:
- `id`: Unique identifier
- `organizationId`: Reference to parent Organization
- `name`: Room name
- `description`: Room description
- `created_date`: Creation timestamp
- `updated_date`: Last update timestamp
- `deleted_date`: Soft delete timestamp

**Status**: ⏳ Pending implementation

### Stage 4: Device CRUD - TO BE IMPLEMENTED
Implement CRUD functionality for the "Device" entity representing equipment that can be deployed or stored.

**Device Categories**: SENSOR, ACTUATOR

**Database Fields**:
- `id`: Unique identifier
- `organizationId`: Reference to parent Organization
- `roomId`: Reference to Room (nullable for storage)
- `guid`: Globally unique identifier
- `inventoryNumber`: Inventory tracking number
- `serialNumber`: Device serial number
- `characteristics`: Device specifications
- `category`: SENSOR or ACTUATOR
- `units`: Measurement units (required for SENSOR)
- `powerConsumption`: Power consumption in watts (required for ACTUATOR)
- `created_date`: Creation timestamp
- `updated_date`: Last update timestamp
- `deleted_date`: Soft delete timestamp

**Requirements**:
- Devices can be deployed to rooms or stored
- ACTUATOR devices must have power consumption specified
- SENSOR devices must have measurement units specified
- System admins can deploy/undeploy devices

**Status**: ⏳ Pending implementation

### Stage 5: Measurement CRUD - TO BE IMPLEMENTED
Implement CRUD functionality for the "Measurement" entity representing sensor data collection.

**Database Fields**:
- `id`: Unique identifier
- `deviceId`: Reference to Device
- `roomId`: Reference to Room
- `value`: Measurement value
- `created_date`: Creation timestamp
- `updated_date`: Last update timestamp
- `deleted_date`: Soft delete timestamp

**Requirements**:
- Accept and store measurement data from deployed sensors
- Admins can filter measurements by:
  - Specific measurement device
  - Time period (day/week/month)

**Status**: ⏳ Pending implementation

### Stage 6: Event CRUD - TO BE IMPLEMENTED
Implement CRUD functionality for the "Event" entity representing actuator on/off events.

**Database Fields**:
- `id`: Unique identifier
- `deviceId`: Reference to Device
- `roomId`: Reference to Room
- `action`: ON/OFF action
- `created_date`: Creation timestamp
- `updated_date`: Last update timestamp
- `deleted_date`: Soft delete timestamp

**Requirements**:
- Accept actuator on/off commands
- Track energy consumption by:
  - Entire enterprise
  - Specific room
  - Time period (day/week/month)

**Status**: ⏳ Pending implementation

### Stage 7: Repository & Documentation - TO BE IMPLEMENTED
- Push source code to GitHub repository
- Provide repository link to instructor
- Submit explanatory notes and documentation

**Status**: ⏳ Pending submission

## Installation & Setup

### Prerequisites
- Go 1.24.0 or later
- PostgreSQL 12 or later
- Git

### Local Development Setup

1. **Clone the repository**
```bash
git clone https://github.com/BohdanBoriak/boilerplate-go-back.git
cd boilerplate-go-back
```

2. **Install Go dependencies**
```bash
go mod download
go mod tidy
```

3. **Set up environment variables**
Create a `.env` file in the project root:
```bash
# Database Configuration
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_NAME=enterprise_system
DATABASE_USER=postgres
DATABASE_PASSWORD=your_secure_password

# Server Configuration
SERVER_PORT=8080
SERVER_HOST=localhost

# JWT Configuration
JWT_SECRET=your_jwt_secret_key_min_32_chars_long
JWT_TTL=3600

# Migration Configuration
MIGRATE_TO_VERSION=latest
MIGRATION_LOCATION=./internal/infra/database/migrations

# File Storage
FILE_STORAGE_LOCATION=./storage
```

4. **Create PostgreSQL database**
```bash
createdb enterprise_system
```

5. **Run database migrations**
The migrations run automatically when the server starts if `MIGRATE_TO_VERSION` is set to "latest".

6. **Start the server**
```bash
go run ./cmd/server/main.go
```

The server will be available at `http://localhost:8080`

### Docker Setup (Optional)

1. **Build Docker image**
```bash
docker build -t enterprise-system:latest .
```

2. **Run with Docker Compose**
```bash
docker-compose up -d
```

## API Documentation

### Authentication Endpoints

**Register**
```
POST /api/v1/auth/register
Content-Type: application/json

{
  "firstName": "John",
  "secondName": "Doe",
  "email": "john@example.com",
  "password": "securePassword123"
}
```

**Login**
```
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "securePassword123"
}

Response:
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": 1,
    "firstName": "John",
    "secondName": "Doe",
    "email": "john@example.com",
    "role": "user"
  }
}
```

**Logout**
```
POST /api/v1/auth/logout
Authorization: Bearer <token>
```

### User Endpoints

**Get Current User**
```
GET /api/v1/users/
Authorization: Bearer <token>
```

**Update User**
```
PUT /api/v1/users/
Authorization: Bearer <token>
Content-Type: application/json

{
  "firstName": "Jane",
  "secondName": "Doe",
  "email": "jane@example.com"
}
```

**Delete User**
```
DELETE /api/v1/users/
Authorization: Bearer <token>
```

### Organization Endpoints (Stage 2)

**Create Organization**
```
POST /api/v1/organizations
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Tech Company Inc",
  "description": "Leading technology solutions provider",
  "city": "Kyiv",
  "address": "123 Innovation Street",
  "lat": 50.4501,
  "lon": 30.5234
}
```

**List Organizations**
```
GET /api/v1/organizations?page=1&count=10
Authorization: Bearer <token>
```

**Get Organization by ID**
```
GET /api/v1/organizations/{id}
Authorization: Bearer <token>
```

**Update Organization**
```
PUT /api/v1/organizations/{id}
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Tech Company Ltd",
  "description": "Updated description",
  "city": "Lviv",
  "address": "456 New Street",
  "lat": 49.8397,
  "lon": 24.0297
}
```

**Delete Organization**
```
DELETE /api/v1/organizations/{id}
Authorization: Bearer <token>
```

## Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go                      # Application entry point
├── config/
│   ├── config.go                        # Configuration management
│   └── container/
│       └── container.go                 # Dependency injection container
├── internal/
│   ├── app/
│   │   ├── auth_service.go             # Authentication service
│   │   ├── user_service.go             # User management service
│   │   ├── organization_service.go     # Organization service
│   │   ├── room_service.go             # Room service (future)
│   │   ├── device_service.go           # Device service (future)
│   │   ├── measurement_service.go      # Measurement service (future)
│   │   └── event_service.go            # Event service (future)
│   ├── domain/
│   │   ├── user.go                     # User entity
│   │   ├── organization.go             # Organization entity
│   │   ├── room.go                     # Room entity (future)
│   │   ├── device.go                   # Device entity (future)
│   │   ├── measurement.go              # Measurement entity (future)
│   │   ├── event.go                    # Event entity (future)
│   │   ├── pagination.go               # Pagination utility
│   │   └── sess.go                     # Session entity
│   └── infra/
│       ├── database/
│       │   ├── user_repository.go      # User data access
│       │   ├── organization_repository.go
│       │   ├── room_repository.go      # Future
│       │   ├── device_repository.go    # Future
│       │   ├── measurement_repository.go # Future
│       │   ├── event_repository.go     # Future
│       │   ├── session_repository.go
│       │   ├── migration.go            # Database migration manager
│       │   └── migrations/             # SQL migration files
│       │       ├── 20221125125907_create_users_table.{up,down}.sql
│       │       ├── 20230104132822_create_sessions_table.{up,down}.sql
│       │       ├── 20260506112820_create_organizations_table.{up,down}.sql
│       │       └── [future migrations]
│       └── http/
│           ├── server.go               # HTTP server setup
│           ├── router.go               # Route definitions
│           ├── controllers/
│           │   ├── auth_controller.go
│           │   ├── user_controller.go
│           │   ├── organization_controller.go
│           │   ├── room_controller.go  # Future
│           │   ├── device_controller.go # Future
│           │   ├── measurement_controller.go # Future
│           │   └── event_controller.go # Future
│           ├── middlewares/
│           │   ├── auth_middleware.go
│           │   └── path_object_middleware.go
│           ├── requests/
│           │   ├── user_request.go
│           │   ├── organization_request.go
│           │   └── validator.go
│           └── resources/
│               ├── user_resource.go
│               └── organization_resource.go
├── go.mod                              # Go module definition
├── go.sum                              # Go module checksums
├── Dockerfile                          # Docker configuration
├── docker-compose.yml                  # Docker Compose setup
├── makefile                            # Build and run targets
├── API_DOCUMENTATION.md                # Detailed API documentation
├── IMPLEMENTATION_STATUS.md            # Current implementation status
└── README.md                           # This file
```

## Database Schema

### Current Tables

**users**
- `id` BIGSERIAL PRIMARY KEY
- `first_name` VARCHAR(255) NOT NULL
- `second_name` VARCHAR(255) NOT NULL
- `email` VARCHAR(255) NOT NULL UNIQUE
- `password` VARCHAR(255) NOT NULL (hashed)
- `role` VARCHAR(50) DEFAULT 'user'
- `created_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
- `updated_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
- `deleted_date` TIMESTAMP (NULL)

**sessions**
- `user_id` BIGINT FOREIGN KEY REFERENCES users(id)
- `uuid` UUID PRIMARY KEY

**organizations**
- `id` BIGSERIAL PRIMARY KEY
- `user_id` BIGINT NOT NULL FOREIGN KEY REFERENCES users(id) ON DELETE CASCADE
- `name` VARCHAR(255) NOT NULL
- `description` TEXT
- `city` VARCHAR(255)
- `address` VARCHAR(255)
- `lat` FLOAT8
- `lon` FLOAT8
- `created_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
- `updated_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
- `deleted_date` TIMESTAMP (NULL)

### Future Tables
- `rooms` - Workspaces within organizations (Stage 3)
- `devices` - Equipment/sensors/actuators (Stage 4)
- `measurements` - Sensor data readings (Stage 5)
- `events` - Actuator on/off events (Stage 6)

## Features

### Security
- JWT token-based authentication
- Password hashing with secure algorithms
- User data isolation
- CORS middleware for cross-origin requests
- Soft delete for data retention

### Architecture
- Clean architecture with separated layers
- Dependency injection for loose coupling
- Repository pattern for data access
- Service layer for business logic
- Middleware pipeline for HTTP request handling

### Database
- PostgreSQL with migrations support
- Soft delete implementation
- Pagination support for list endpoints
- Foreign key constraints and cascading deletes
- Indexes for optimized queries

## API Testing

### Using Postman

1. Import the API collection from the project
2. Set the base URL to `http://localhost:8080`
3. Obtain JWT token via login endpoint
4. Use token in Authorization header: `Bearer <token>`
5. Test all endpoints

### Example Postman Environment Variables
```json
{
  "base_url": "http://localhost:8080",
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user_id": 1,
  "organization_id": 1
}
```

## Development Guidelines

### Code Style
- Follow Go naming conventions (CamelCase for exported, camelCase for unexported)
- Use meaningful variable and function names
- Add comments for exported functions and types
- Run `go fmt` before committing

### Testing
```bash
go test ./...
go test -v ./...
go test -cover ./...
```

### Building
```bash
go build -o server ./cmd/server
```

### Running Linter
```bash
golangci-lint run
```

## Common Issues

### Database Connection Error
```
Unable to create new DB session: connection refused
```
**Solution**: Ensure PostgreSQL is running and credentials are correct in `.env` file

### Migration Errors
```
Unable to read migration version
```
**Solution**: Check migration files exist in `MIGRATION_LOCATION` path

### Port Already in Use
```
listen tcp :8080: bind: address already in use
```
**Solution**: Change `SERVER_PORT` in `.env` or kill process using port 8080

## License

This project is part of a university coursework assignment.

## Submission Requirements

- [ ] All code pushed to GitHub
- [ ] README with setup instructions
- [ ] API documentation
- [ ] Explanatory notes (Пояснювальна записка)
- [ ] Implementation status document
- [ ] GitHub repository link provided to instructor