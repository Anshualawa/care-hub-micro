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

type Doctor struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Role         string   `json:"role"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	Department   string   `json:"department"`
	Specialization string `json:"specialization"`
	Bio          string   `json:"bio"`
	Education    []string `json:"education"`
	Experience   []string `json:"experience"`
	ProfileImage string   `json:"profileImage"`
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
	Name: "CareHub Hospital",
	Address: "789 Health Ave, Metropolis, USA",
	Email: "info@carehub.com",
	Phone: "555-111-2222",
	Description: "Modern hospital with the best medical team in the region.",
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

// --- Doctor Handlers ---
func getDoctors(c *gin.Context) {
	c.JSON(http.StatusOK, doctors)
}

func getDoctor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for _, d := range doctors {
		if d.ID == id {
			c.JSON(http.StatusOK, d)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
}

func updateDoctor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var updated Doctor
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, d := range doctors {
		if d.ID == id {
			updated.ID = id
			doctors[i] = updated
			c.JSON(http.StatusOK, updated)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
}

// --- Blog Handlers ---
func getBlogs(c *gin.Context) {
	c.JSON(http.StatusOK, blogs)
}

func getBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for _, b := range blogs {
		if b.ID == id {
			c.JSON(http.StatusOK, b)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
}

func createBlog(c *gin.Context) {
	var newBlog Blog
	if err := c.ShouldBindJSON(&newBlog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newBlog.ID = len(blogs) + 1
	blogs = append(blogs, newBlog)
	c.JSON(http.StatusCreated, newBlog)
}

func updateBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var updated Blog
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, b := range blogs {
		if b.ID == id {
			updated.ID = id
			blogs[i] = updated
			c.JSON(http.StatusOK, updated)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
}

func deleteBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for i, b := range blogs {
		if b.ID == id {
			blogs = append(blogs[:i], blogs[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Blog deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
}

// --- Intern Handlers ---
func getInterns(c *gin.Context) {
	c.JSON(http.StatusOK, interns)
}

func getIntern(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for _, i := range interns {
		if i.ID == id {
			c.JSON(http.StatusOK, i)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Intern not found"})
}

// --- Hospital Handler ---
func getHospital(c *gin.Context) {
	c.JSON(http.StatusOK, hospital)
}
