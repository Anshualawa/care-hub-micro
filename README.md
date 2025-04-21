
# CareHub Healthcare Microservice

A modern healthcare management application built with Go (Gin) and React.

## Features

- **Patient Management:** Create, view, update, and delete patient records
- **Appointment Scheduling:** Manage appointments with calendar integration
- **Health Metrics Tracking:** Monitor patient health data with visual charts
- **User Authentication:** Secure login system with role-based access

## Technologies Used

### Backend
- Go with Gin web framework
- RESTful API architecture
- In-memory data storage (can be extended to use a database)

### Frontend
- React with TypeScript
- TailwindCSS for styling
- Recharts for data visualization
- React Query for API data management

## Getting Started

### Running the Backend

```bash
# Navigate to the api directory
cd api

# Install dependencies
go mod tidy

# Run the service
go run main.go
```

The backend service will run on port 8090.

### Running the Frontend

```bash
# In the project root directory
npm install
npm run dev
```

The frontend application will run on port 8080.

## API Documentation

See the [API README](api/README.md) for detailed endpoint documentation.

## Demo Credentials

- Admin: username: `admin`, password: `admin123`
- Doctor: username: `doctor`, password: `doctor123`

## Project Structure

```
/
├── api/                # Go backend
│   ├── main.go         # Main application entry
│   └── go.mod          # Go module definition
│
├── src/
│   ├── components/     # UI Components
│   ├── context/        # React Context providers  
│   ├── pages/          # Application pages
│   ├── services/       # API service layer
│   └── utils/          # Utility functions
│
└── public/             # Static assets
```

## Next Steps / Future Enhancements

- Add persistent database storage
- Implement JWT authentication with refresh tokens
- Add more health metrics and visualization options
- Add reporting and export functionality
- Mobile application support
