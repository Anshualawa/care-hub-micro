package main

import (
	"carehub-microservice/db"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	ID         int       `json:"id"`
	PatientID  int       `json:"patientId"`
	Type       string    `json:"type"`
	Value      float64   `json:"value"`
	Unit       string    `json:"unit"`
	RecordedAt time.Time `json:"recordedAt"`
}

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Department string `json:"department"`
}

type Doctor struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Role           string   `json:"role"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	Department     string   `json:"department"`
	Specialization string   `json:"specialization"`
	Bio            string   `json:"bio"`
	Education      []string `json:"education"`
	Experience     []string `json:"experience"`
	ProfileImage   string   `json:"profileImage"`
}

type Blog struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Excerpt     string   `json:"excerpt"`
	CoverImage  string   `json:"coverImage"`
	AuthorId    int      `json:"authorId"`
	AuthorName  string   `json:"authorName"`
	PublishedAt string   `json:"publishedAt"`
	UpdatedAt   string   `json:"updatedAt"`
	Tags        []string `json:"tags"`
}

type Intern struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
}

type Hospital struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
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

var doctors = []Doctor{
	{ID: 1, Name: "Dr. Jane Smith", Role: "doctor", Email: "jane.smith@carehub.com", Phone: "555-123-4567", Department: "Cardiology", Specialization: "Cardiology", Bio: "Dr. Smith is a board-certified cardiologist with over 15 years of experience.", Education: []string{"MD, Harvard Medical School", "Residency, Mayo Clinic"}, Experience: []string{"Senior Cardiologist, Mayo Clinic (2015-2020)", "Chief of Cardiology, CareHub Hospital (2020-Present)"}, ProfileImage: ""},
	{ID: 2, Name: "Dr. Robert Chen", Role: "doctor", Email: "robert.chen@carehub.com", Phone: "555-234-5678", Department: "Neurology", Specialization: "Neurology", Bio: "Dr. Chen specializes in neurological disorders.", Education: []string{"MD, Johns Hopkins University", "Fellowship, Cleveland Clinic"}, Experience: []string{"Neurologist, Cleveland Clinic (2013-2018)", "Senior Neurologist, CareHub Hospital (2018-Present)"}, ProfileImage: "https://randomuser.me/api/portraits/men/32.jpg"},
	{ID: 3, Name: "Dr. Maria Rodriguez", Role: "doctor", Email: "maria.rodriguez@carehub.com", Phone: "555-345-6789", Department: "Pediatrics", Specialization: "Pediatrics", Bio: "Dr. Rodriguez has dedicated her career to children's health.", Education: []string{"MD, Stanford University", "Residency, Children's Hospital of Philadelphia"}, Experience: []string{"Pediatrician, Boston Children's Hospital (2016-2021)", "Lead Pediatrician, CareHub Hospital (2021-Present)"}, ProfileImage: "https://randomuser.me/api/portraits/women/45.jpg"},
	{ID: 4, Name: "Dr. James Wilson", Role: "doctor", Email: "james.wilson@carehub.com", Phone: "555-456-7890", Department: "Orthopedics", Specialization: "Orthopedic Surgery", Bio: "Dr. Wilson is an orthopedic surgeon specializing in sports injuries.", Education: []string{"MD, University of Michigan", "Orthopedic Fellowship, Hospital for Special Surgery"}, Experience: []string{"Orthopedic Surgeon, UCSF Medical Center (2012-2019)", "Chief of Orthopedics, CareHub Hospital (2019-Present)"}, ProfileImage: ""},
}

var blogs = []Blog{
	{ID: 1, Title: "Heart Health Tips", Content: "Eat well, exercise, and manage stress.", Excerpt: "Stay heart-healthy!", CoverImage: "", AuthorId: 1, AuthorName: "Dr. Jane Smith", PublishedAt: "2024-06-11T11:00:00Z", UpdatedAt: "2024-06-11T14:00:00Z", Tags: []string{"cardiology"}},
	{ID: 2, Title: "Understanding Migraines", Content: "Migraines affect millions. Here's what helps.", Excerpt: "Guide to controlling migraines.", CoverImage: "", AuthorId: 2, AuthorName: "Dr. Robert Chen", PublishedAt: "2024-05-21T09:30:00Z", UpdatedAt: "", Tags: []string{"neurology"}},
}

var interns = []Intern{
	{ID: 1, Name: "Alex Green", Email: "alex.green@carehub.com", Department: "Cardiology"},
	{ID: 2, Name: "Priya Patel", Email: "priya.patel@carehub.com", Department: "Pediatrics"},
}

var hospital = Hospital{
	Name:        "CareHub Hospital",
	Address:     "789 Health Ave, Metropolis, USA",
	Email:       "info@carehub.com",
	Phone:       "555-111-2222",
	Description: "Modern hospital with the best medical team in the region.",
}

func main() {
	// Initialize database connection
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.CloseDB()

	r := gin.Default()

	// Configure CORS to allow frontend requests
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://192.168.1.7:8080"},
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
		auth.POST("/signup", signup)
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

		// Doctor endpoints
		api.GET("/doctors", getDoctors)
		api.GET("/doctors/:id", getDoctor)
		api.PUT("/doctors/:id", updateDoctor)

		// Blog endpoints
		api.GET("/blogs", getBlogs)
		api.GET("/blogs/:id", getBlog)
		api.POST("/blogs", createBlog)
		api.PUT("/blogs/:id", updateBlog)
		api.DELETE("/blogs/:id", deleteBlog)

		// Intern endpoints
		api.GET("/interns", getInterns)
		api.GET("/interns/:id", getIntern)

		// Hospital endpoint (only GET)
		api.GET("/hospital", getHospital)
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

	// Query user from database
	var user User
	query := "SELECT id, username, password, role, name, email, phone, department FROM users WHERE username = ? LIMIT 1"
	err := db.DB.QueryRow(query, loginData.Username).Scan(
		&user.ID, &user.Username, &user.Password, &user.Role, &user.Name, &user.Email, &user.Phone, &user.Department)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Simple authentication (in real app, use proper password hashing)
	if user.Password == loginData.Password {
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

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

func signup(c *gin.Context) {
	var userData struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		Role       string `json:"role"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		Phone      string `json:"phone"`
		Department string `json:"department"`
	}

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid signup data"})
		return
	}

	// Check if username already exists
	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", userData.Username).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// Check if email already exists
	err = db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", userData.Email).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// Insert new user
	query := `INSERT INTO users (username, password, role, name, email, phone, department) 
			  VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := db.DB.Exec(query,
		userData.Username, userData.Password, userData.Role,
		userData.Name, userData.Email, userData.Phone, userData.Department)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"userId":  id,
	})
}

// Patient Handlers
func getPatients(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, first_name, last_name, date_of_birth, email, phone, address, created_at FROM patients")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var patients []Patient
	for rows.Next() {
		var patient Patient
		var dbTime time.Time
		err := rows.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.DateOfBirth,
			&patient.Email, &patient.Phone, &patient.Address, &dbTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Row scan error"})
			return
		}
		patient.CreatedAt = dbTime
		patients = append(patients, patient)
	}

	c.JSON(http.StatusOK, patients)
}

func getPatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var patient Patient
	var dbTime time.Time
	query := "SELECT id, first_name, last_name, date_of_birth, email, phone, address, created_at FROM patients WHERE id = ?"
	err = db.DB.QueryRow(query, id).Scan(
		&patient.ID, &patient.FirstName, &patient.LastName, &patient.DateOfBirth,
		&patient.Email, &patient.Phone, &patient.Address, &dbTime)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	patient.CreatedAt = dbTime
	c.JSON(http.StatusOK, patient)
}

func createPatient(c *gin.Context) {
	var newPatient Patient
	if err := c.ShouldBindJSON(&newPatient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO patients (first_name, last_name, date_of_birth, email, phone, address) 
			  VALUES (?, ?, ?, ?, ?, ?)`
	result, err := db.DB.Exec(query,
		newPatient.FirstName, newPatient.LastName, newPatient.DateOfBirth,
		newPatient.Email, newPatient.Phone, newPatient.Address)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient: " + err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get inserted ID"})
		return
	}

	newPatient.ID = int(id)
	newPatient.CreatedAt = time.Now()

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

	query := `UPDATE patients SET first_name = ?, last_name = ?, date_of_birth = ?, 
			 email = ?, phone = ?, address = ? WHERE id = ?`
	result, err := db.DB.Exec(query,
		updatedPatient.FirstName, updatedPatient.LastName, updatedPatient.DateOfBirth,
		updatedPatient.Email, updatedPatient.Phone, updatedPatient.Address, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient: " + err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get affected rows"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// Fetch the updated patient with created_at timestamp
	var dbTime time.Time
	queryGet := "SELECT created_at FROM patients WHERE id = ?"
	err = db.DB.QueryRow(queryGet, id).Scan(&dbTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated patient"})
		return
	}

	updatedPatient.ID = id
	updatedPatient.CreatedAt = dbTime

	c.JSON(http.StatusOK, updatedPatient)
}

func deletePatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	result, err := db.DB.Exec("DELETE FROM patients WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete patient: " + err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get affected rows"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted"})
}

// Appointment Handlers
func getAppointments(c *gin.Context) {
	patientID := c.Query("patientId")

	var rows *sql.Rows
	var err error

	if patientID != "" {
		rows, err = db.DB.Query("SELECT id, patient_id, date_time, description, status, doctor FROM appointments WHERE patient_id = ?", patientID)
	} else {
		rows, err = db.DB.Query("SELECT id, patient_id, date_time, description, status, doctor FROM appointments")
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var appointments []Appointment
	for rows.Next() {
		var appointment Appointment
		var dbTime time.Time
		err := rows.Scan(&appointment.ID, &appointment.PatientID, &dbTime,
			&appointment.Description, &appointment.Status, &appointment.Doctor)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Row scan error"})
			return
		}
		appointment.DateTime = dbTime
		appointments = append(appointments, appointment)
	}

	c.JSON(http.StatusOK, appointments)
}

func getAppointment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var appointment Appointment
	var dbTime time.Time
	query := "SELECT id, patient_id, date_time, description, status, doctor FROM appointments WHERE id = ?"
	err = db.DB.QueryRow(query, id).Scan(
		&appointment.ID, &appointment.PatientID, &dbTime,
		&appointment.Description, &appointment.Status, &appointment.Doctor)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + err.Error()})
		}
		return
	}

	appointment.DateTime = dbTime
	c.JSON(http.StatusOK, appointment)
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

	query := `INSERT INTO appointments (patient_id, date_time, description, status, doctor) 
			  VALUES (?, ?, ?, ?, ?)`
	result, err := db.DB.Exec(query,
		appointmentData.PatientID, dateTime, appointmentData.Description,
		appointmentData.Status, appointmentData.Doctor)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment: " + err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get inserted ID"})
		return
	}

	newAppointment := Appointment{
		ID:          int(id),
		PatientID:   appointmentData.PatientID,
		DateTime:    dateTime,
		Description: appointmentData.Description,
		Status:      appointmentData.Status,
		Doctor:      appointmentData.Doctor,
	}

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
		DateTime    string `json:"dateTime"`
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
		dateTime, err = time.Parse("2006-01-02T15:04:05", appointmentData.DateTime)
		if err != nil {
			dateTime, err = time.Parse(time.RFC3339, appointmentData.DateTime)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid datetime format: " + err.Error()})
				return
			}
		}
	}

	query := `UPDATE appointments SET patient_id = ?, date_time = ?, description = ?,
			 status = ?, doctor = ? WHERE id = ?`
	result, err := db.DB.Exec(query,
		appointmentData.PatientID, dateTime, appointmentData.Description,
		appointmentData.Status, appointmentData.Doctor, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update appointment: " + err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get affected rows"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	updatedAppointment := Appointment{
		ID:          id,
		PatientID:   appointmentData.PatientID,
		DateTime:    dateTime,
		Description: appointmentData.Description,
		Status:      appointmentData.Status,
		Doctor:      appointmentData.Doctor,
	}

	c.JSON(http.StatusOK, updatedAppointment)
}

func deleteAppointment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	result, err := db.DB.Exec("DELETE FROM appointments WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete appointment: " + err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get affected rows"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment deleted"})
}

// Health Metric Handlers
func getPatientHealthMetrics(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID format"})
		return
	}

	rows, err := db.DB.Query("SELECT id, patient_id, type, value, unit, recorded_at FROM health_metrics WHERE patient_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var metrics []HealthMetric
	for rows.Next() {
		var metric HealthMetric
		var dbTime time.Time
		err := rows.Scan(&metric.ID, &metric.PatientID, &metric.Type,
			&metric.Value, &metric.Unit, &dbTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Row scan error"})
			return
		}
		metric.RecordedAt = dbTime
		metrics = append(metrics, metric)
	}

	c.JSON(http.StatusOK, metrics)
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

	query := `INSERT INTO health_metrics (patient_id, type, value, unit) 
			  VALUES (?, ?, ?, ?)`
	result, err := db.DB.Exec(query, patientID, newMetric.Type, newMetric.Value, newMetric.Unit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record health metric: " + err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get inserted ID"})
		return
	}

	// Fetch the created health metric with recorded_at timestamp
	var dbTime time.Time
	queryGet := "SELECT recorded_at FROM health_metrics WHERE id = ?"
	err = db.DB.QueryRow(queryGet, id).Scan(&dbTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve recorded metric"})
		return
	}

	newMetric.ID = int(id)
	newMetric.PatientID = patientID
	newMetric.RecordedAt = dbTime

	c.JSON(http.StatusCreated, newMetric)
}

// --- Doctor Handlers ---
func getDoctors(c *gin.Context) {
	rows, err := db.DB.Query(`SELECT id, name, role, email, phone, department, 
		specialization, bio, profile_image FROM doctors`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var doctors []Doctor
	for rows.Next() {
		var doctor Doctor
		err := rows.Scan(&doctor.ID, &doctor.Name, &doctor.Role, &doctor.Email,
			&doctor.Phone, &doctor.Department, &doctor.Specialization, &doctor.Bio, &doctor.ProfileImage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Row scan error"})
			return
		}

		// Get education and experience for each doctor
		doctor.Education = []string{}
		doctor.Experience = []string{}

		// In a real implementation, we'd fetch these from related tables
		// For simplicity, we'll use placeholder data
		if doctor.ID == 1 {
			doctor.Education = []string{"MD, Harvard Medical School", "Residency, Mayo Clinic"}
			doctor.Experience = []string{"Senior Cardiologist, Mayo Clinic (2015-2020)", "Chief of Cardiology, CareHub Hospital (2020-Present)"}
		} else if doctor.ID == 2 {
			doctor.Education = []string{"MD, Johns Hopkins University", "Fellowship, Cleveland Clinic"}
			doctor.Experience = []string{"Neurologist, Cleveland Clinic (2013-2018)", "Senior Neurologist, CareHub Hospital (2018-Present)"}
		}

		doctors = append(doctors, doctor)
	}

	c.JSON(http.StatusOK, doctors)
}

func getDoctor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var doctor Doctor
	query := `SELECT id, name, role, email, phone, department, 
		specialization, bio, profile_image FROM doctors WHERE id = ?`
	err = db.DB.QueryRow(query, id).Scan(&doctor.ID, &doctor.Name, &doctor.Role,
		&doctor.Email, &doctor.Phone, &doctor.Department, &doctor.Specialization,
		&doctor.Bio, &doctor.ProfileImage)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	// Get education and experience (simplified)
	if doctor.ID == 1 {
		doctor.Education = []string{"MD, Harvard Medical School", "Residency, Mayo Clinic"}
		doctor.Experience = []string{"Senior Cardiologist, Mayo Clinic (2015-2020)", "Chief of Cardiology, CareHub Hospital (2020-Present)"}
	} else if doctor.ID == 2 {
		doctor.Education = []string{"MD, Johns Hopkins University", "Fellowship, Cleveland Clinic"}
		doctor.Experience = []string{"Neurologist, Cleveland Clinic (2013-2018)", "Senior Neurologist, CareHub Hospital (2018-Present)"}
	}

	c.JSON(http.StatusOK, doctor)
}

func updateDoctor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var doctor Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update basic doctor information
	query := `UPDATE doctors SET name = ?, role = ?, email = ?, phone = ?,
		department = ?, specialization = ?, bio = ?, profile_image = ? WHERE id = ?`
	result, err := db.DB.Exec(query, doctor.Name, doctor.Role, doctor.Email, doctor.Phone,
		doctor.Department, doctor.Specialization, doctor.Bio, doctor.ProfileImage, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update doctor: " + err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get affected rows"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}

	// In a real implementation, we'd update education and experience in related tables

	doctor.ID = id
	c.JSON(http.StatusOK, doctor)
}

// --- Blog Handlers ---
func getBlogs(c *gin.Context) {
	rows, err := db.DB.Query(`SELECT id, title, content, excerpt, cover_image, 
		author_id, author_name, published_at, updated_at FROM blogs`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var blogs []Blog
	for rows.Next() {
		var blog Blog
		var publishedAt, updatedAt string
		err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Excerpt,
			&blog.CoverImage, &blog.AuthorId, &blog.AuthorName, &publishedAt, &updatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Row scan error"})
			return
		}

		blog.PublishedAt = publishedAt
		blog.UpdatedAt = updatedAt

		// Get tags for each blog (simplification)
		if blog.ID == 1 {
			blog.Tags = []string{"cardiology"}
		} else if blog.ID == 2 {
			blog.Tags = []string{"neurology"}
		}

		blogs = append(blogs, blog)
	}

	c.JSON(http.StatusOK, blogs)
}

func getBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var blog Blog
	var publishedAt, updatedAt string
	query := `SELECT id, title, content, excerpt, cover_image, 
		author_id, author_name, published_at, updated_at FROM blogs WHERE id = ?`
	err = db.DB.QueryRow(query, id).Scan(&blog.ID, &blog.Title, &blog.Content,
		&blog.Excerpt, &blog.CoverImage, &blog.AuthorId, &blog.AuthorName, &publishedAt, &updatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	blog.PublishedAt = publishedAt
	blog.UpdatedAt = updatedAt

	// Get tags (simplified)
	if blog.ID == 1 {
		blog.Tags = []string{"cardiology"}
	} else if blog.ID == 2 {
		blog.Tags = []string{"neurology"}
	}

	c.JSON(http.StatusOK, blog)
}

func createBlog(c *gin.Context) {
	var blog Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().Format(time.RFC3339)

	query := `INSERT INTO blogs (title, content, excerpt, cover_image, author_id, author_name, published_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := db.DB.Exec(query, blog.Title, blog.Content, blog.Excerpt,
		blog.CoverImage, blog.AuthorId, blog.AuthorName, now)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog: " + err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get inserted ID"})
		return
	}

	blog.ID = int(id)
	blog.PublishedAt = now

	// In a real implementation, we would save tags to a related table

	c.JSON(http.StatusCreated, blog)
}

func updateBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var blog Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().Format(time.RFC3339)

	query := `UPDATE blogs SET title = ?, content = ?, excerpt = ?, cover_image = ?, 
		author_id = ?, author_name = ?, updated_at = ? WHERE id = ?`
	result, err := db.DB.Exec(query, blog.Title, blog.Content, blog.Excerpt, blog.CoverImage,
		blog.AuthorId, blog.AuthorName, now, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog: " + err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get affected rows"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	// Fetch the updated blog with timestamps
	var publishedAt string
	queryGet := "SELECT published_at FROM blogs WHERE id = ?"
	err = db.DB.QueryRow(queryGet, id).Scan(&publishedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated blog"})
		return
	}

	blog.ID = id
	blog.PublishedAt = publishedAt
	blog.UpdatedAt = now

	c.JSON(http.StatusOK, blog)
}

func deleteBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := db.DB.Exec("DELETE FROM blogs WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog: " + err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get affected rows"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted"})
}

// --- Intern Handlers ---
func getInterns(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, name, email, department FROM interns")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var interns []Intern
	for rows.Next() {
		var intern Intern
		err := rows.Scan(&intern.ID, &intern.Name, &intern.Email, &intern.Department)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Row scan error"})
			return
		}
		interns = append(interns, intern)
	}

	c.JSON(http.StatusOK, interns)
}

func getIntern(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var intern Intern
	query := "SELECT id, name, email, department FROM interns WHERE id = ?"
	err = db.DB.QueryRow(query, id).Scan(&intern.ID, &intern.Name, &intern.Email, &intern.Department)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Intern not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	c.JSON(http.StatusOK, intern)
}

// --- Hospital Handler ---
func getHospital(c *gin.Context) {
	// For this demo, we'll use a hardcoded hospital since it's a singleton
	hospital := Hospital{
		Name:        "CareHub Hospital",
		Address:     "789 Health Ave, Metropolis, USA",
		Email:       "info@carehub.com",
		Phone:       "555-111-2222",
		Description: "Modern hospital with the best medical team in the region.",
	}

	c.JSON(http.StatusOK, hospital)
}
