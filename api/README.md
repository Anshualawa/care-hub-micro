
# Healthcare Microservice API

This is a Go-based microservice for healthcare management built with the Gin web framework.

## Features

- Patient management (CRUD operations)
- Appointment scheduling and management
- Health metrics tracking
- Basic authentication

## Running the Service

```bash
# Navigate to the api directory
cd api

# Install dependencies
go mod tidy

# Run the service
go run main.go
```

The service will start on port 8090.

## API Endpoints

### Authentication

- `POST /auth/login` - Login with username and password

### Patients

- `GET /api/patients` - Get all patients
- `GET /api/patients/:id` - Get a specific patient
- `POST /api/patients` - Create a new patient
- `PUT /api/patients/:id` - Update a patient
- `DELETE /api/patients/:id` - Delete a patient

### Appointments

- `GET /api/appointments` - Get all appointments
- `GET /api/appointments?patientId=1` - Get appointments for a specific patient
- `GET /api/appointments/:id` - Get a specific appointment
- `POST /api/appointments` - Create a new appointment
- `PUT /api/appointments/:id` - Update an appointment
- `DELETE /api/appointments/:id` - Delete an appointment

### Health Metrics

- `GET /api/patients/:id/metrics` - Get health metrics for a specific patient
- `POST /api/patients/:id/metrics` - Record a health metric for a patient

## Demo Users

- Admin: username: `admin`, password: `admin123`
- Doctor: username: `doctor`, password: `doctor123`
