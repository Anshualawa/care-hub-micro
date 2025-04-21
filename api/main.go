package main

import (
	"net/http"
	"strconv"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

// Models
type Patient struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DateOfBirth string    `json:"dateOfBirth"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Appointment struct {
	ID          int       `json:"id"`
	PatientID   int       `json:"patientId"`
	DateTime    time.Time `json:"dateTime"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Doctor      string    `json:"doctor"`
}

type HealthMetric struct {
	ID        int       `json:"id"`
	PatientID int       `json:"patientId"`
	Type      string    `json:"type"`
	Value     float64   `json:"value"`
	Unit      string    `json:"unit"`
	RecordedAt time.Time `json:"recordedAt"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Department string `json:"department"`
}

// Mock data stores (would be replaced by a database in production)
var patients = []Patient{
	{ID: 1, FirstName: "John", LastName: "Doe", DateOfBirth: "1980-05-15", Email: "john.doe@example.com", Phone: "555-123-4567", Address: "123 Main St, Anytown, CA", CreatedAt: time.Now().Add(-24 * time.Hour)},
	{ID: 2, FirstName: "Jane", LastName: "Smith", DateOfBirth: "1975-08-21", Email: "jane.smith@example.com", Phone: "555-987-6543", Address: "456 Oak Ave, Somewhere, NY", CreatedAt: time.Now().Add(-48 * time.Hour)},
}

var appointments = []Appointment{
	{ID: 1, PatientID: 1, DateTime: time.Now().Add(48 * time.Hour), Description: "Annual checkup", Status: "Scheduled", Doctor: "Dr. Brown"},
	{ID: 2, PatientID: 2, DateTime: time.Now().Add(72 * time.Hour), Description: "Follow-up", Status: "Scheduled", Doctor: "Dr. Johnson"},
}

var healthMetrics = []HealthMetric{
	{ID: 1, PatientID: 1, Type: "Blood Pressure", Value: 120.80, Unit: "mmHg", RecordedAt: time.Now().Add(-24 * time.Hour)},
	{ID: 2, PatientID: 1, Type: "Heart Rate", Value: 72, Unit: "bpm", RecordedAt: time.Now().Add(-24 * time.Hour)},
	{ID: 3, PatientID: 2, Type: "Blood Pressure", Value: 118.75, Unit: "mmHg", RecordedAt: time.Now().Add(-48 * time.Hour)},
}

var users = []User{
	{ID: 1, Username: "admin", Password: "admin123", Role: "admin", Name: "Administrator", Email: "admin@carehub.com", Phone: "555-100-0001", Department: "Administration"},
	{ID: 2, Username: "doctor", Password: "doctor123", Role: "doctor", Name: "Dr. Sarah Williams", Email: "sarah.williams@carehub.com", Phone: "555-100-0002", Department: "Cardiology"},
	{ID: 3, Username: "superadmin", Password: "super123", Role: "superadmin", Name: "System Administrator", Email: "sysadmin@carehub.com", Phone: "555-100-0003", Department: "IT"},
	{ID: 4, Username: "nurse", Password: "nurse123", Role: "nurse", Name: "Nancy White", Email: "nancy.white@carehub.com", Phone: "555-100-0004", Department: "Emergency"},
	{ID: 5, Username: "intern", Password: "intern123", Role: "intern", Name: "Dr. Michael Lee", Email: "michael.lee@carehub.com", Phone: "555-100-0005", Department: "Pediatrics"},
	{ID: 6, Username: "patient", Password: "patient123", Role: "patient", Name: "John Doe", Email: "john.doe@example.com", Phone: "555-123-4567", Department: ""},
}

func main() {
	r := gin.Default()

	// Configure CORS to allow frontend requests
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Authentication endpoints
	auth := r.Group("/auth")
	{
		auth.POST("/login", login)
	}

	// API routes
	api := r.Group("/api")
	{
		// Patient endpoints
		api.GET("/patients", getPatients)
		api.GET("/patients/:id", getPatient)
		api.POST("/patients", createPatient)
		api.PUT("/patients/:id", updatePatient)
		api.DELETE("/patients/:id", deletePatient)

		// Appointment endpoints
		api.GET("/appointments", getAppointments)
		api.GET("/appointments/:id", getAppointment)
		api.POST("/appointments", createAppointment)
		api.PUT("/appointments/:id", updateAppointment)
		api.DELETE("/appointments/:id", deleteAppointment)

		// Health metric endpoints
		api.GET("/patients/:id/metrics", getPatientHealthMetrics)
		api.POST("/patients/:id/metrics", recordHealthMetric)
	}

	// Server setup
	fmt.Println("Starting server on port 8090...")
	r.Run(":8090")
}

// Authentication Handlers
func login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login data"})
		return
	}

	// Simple authentication (in real app, use proper password hashing)
	for _, user := range users {
		if user.Username == loginData.Username && user.Password == loginData.Password {
			c.JSON(http.StatusOK, gin.H{
				"token": "sample-jwt-token", // In real app, generate JWT token
				"user": gin.H{
					"id":   user.ID,
					"name": user.Name,
					"role": user.Role,
				},
			})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

// Patient Handlers
func getPatients(c *gin.Context) {
	c.JSON(http.StatusOK, patients)
}

func getPatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	for _, patient := range patients {
		if patient.ID == id {
			c.JSON(http.StatusOK, patient)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
}

func createPatient(c *gin.Context) {
	var newPatient Patient
	if err := c.ShouldBindJSON(&newPatient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate new ID (in a real app, database would handle this)
	newPatient.ID = len(patients) + 1
	newPatient.CreatedAt = time.Now()

	patients = append(patients, newPatient)
	c.JSON(http.StatusCreated, newPatient)
}

func updatePatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updatedPatient Patient
	if err := c.ShouldBindJSON(&updatedPatient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, patient := range patients {
		if patient.ID == id {
			updatedPatient.ID = id
			updatedPatient.CreatedAt = patient.CreatedAt
			patients[i] = updatedPatient
			c.JSON(http.StatusOK, updatedPatient)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
}

func deletePatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	for i, patient := range patients {
		if patient.ID == id {
			patients = append(patients[:i], patients[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Patient deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
}

// Appointment Handlers
func getAppointments(c *gin.Context) {
	patientID := c.Query("patientId")
	if patientID != "" {
		id, err := strconv.Atoi(patientID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID format"})
			return
		}

		var patientAppointments []Appointment
		for _, appointment := range appointments {
			if appointment.PatientID == id {
				patientAppointments = append(patientAppointments, appointment)
			}
		}
		c.JSON(http.StatusOK, patientAppointments)
		return
	}

	c.JSON(http.StatusOK, appointments)
}

func getAppointment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	for _, appointment := range appointments {
		if appointment.ID == id {
			c.JSON(http.StatusOK, appointment)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
}

func createAppointment(c *gin.Context) {
	var appointmentData struct {
		PatientID   int    `json:"patientId"`
		DateTime    string `json:"dateTime"` // Accept as string initially
		Description string `json:"description"`
		Status      string `json:"status"`
		Doctor      string `json:"doctor"`
	}

	if err := c.ShouldBindJSON(&appointmentData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse the datetime string into a time.Time
	dateTime, err := time.Parse("2006-01-02T15:04", appointmentData.DateTime)
	if err != nil {
		// If that fails, try with seconds
		dateTime, err = time.Parse("2006-01-02T15:04:05", appointmentData.DateTime)
		if err != nil {
			// If that fails too, try RFC3339 format
			dateTime, err = time.Parse(time.RFC3339, appointmentData.DateTime)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid datetime format: " + err.Error()})
				return
			}
		}
	}

	// Generate new ID
	newAppointment := Appointment{
		ID:          len(appointments) + 1,
		PatientID:   appointmentData.PatientID,
		DateTime:    dateTime,
		Description: appointmentData.Description,
		Status:      appointmentData.Status,
		Doctor:      appointmentData.Doctor,
	}

	appointments = append(appointments, newAppointment)
	c.JSON(http.StatusCreated, newAppointment)
}

func updateAppointment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var appointmentData struct {
		PatientID   int    `json:"patientId"`
		DateTime    string `json:"dateTime"` // Accept as string initially
		Description string `json:"description"`
		Status      string `json:"status"`
		Doctor      string `json:"doctor"`
	}

	if err := c.ShouldBindJSON(&appointmentData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse the datetime string into a time.Time
	dateTime, err := time.Parse("2006-01-02T15:04", appointmentData.DateTime)
	if err != nil {
		// If that fails, try with seconds
		dateTime, err = time.Parse("2006-01-02T15:04:05", appointmentData.DateTime)
		if err != nil {
			// If that fails too, try RFC3339 format
			dateTime, err = time.Parse(time.RFC3339, appointmentData.DateTime)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid datetime format: " + err.Error()})
				return
			}
		}
	}

	for i, appointment := range appointments {
		if appointment.ID == id {
			appointments[i].PatientID = appointmentData.PatientID
			appointments[i].DateTime = dateTime
			appointments[i].Description = appointmentData.Description
			appointments[i].Status = appointmentData.Status
			appointments[i].Doctor = appointmentData.Doctor
			c.JSON(http.StatusOK, appointments[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
}

func deleteAppointment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	for i, appointment := range appointments {
		if appointment.ID == id {
			appointments = append(appointments[:i], appointments[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Appointment deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
}

// Health Metric Handlers
func getPatientHealthMetrics(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID format"})
		return
	}

	var patientMetrics []HealthMetric
	for _, metric := range healthMetrics {
		if metric.PatientID == id {
			patientMetrics = append(patientMetrics, metric)
		}
	}
	c.JSON(http.StatusOK, patientMetrics)
}

func recordHealthMetric(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID format"})
		return
	}

	var newMetric HealthMetric
	if err := c.ShouldBindJSON(&newMetric); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate new ID
	newMetric.ID = len(healthMetrics) + 1
	newMetric.PatientID = patientID
	newMetric.RecordedAt = time.Now()

	healthMetrics = append(healthMetrics, newMetric)
	c.JSON(http.StatusCreated, newMetric)
}
