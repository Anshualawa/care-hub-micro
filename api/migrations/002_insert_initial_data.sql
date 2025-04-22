
-- Insert Users
INSERT INTO users (username, password, role, name, email, phone, department) VALUES
('admin', 'admin123', 'admin', 'Administrator', 'admin@carehub.com', '555-100-0001', 'Administration'),
('doctor', 'doctor123', 'doctor', 'Dr. Sarah Williams', 'sarah.williams@carehub.com', '555-100-0002', 'Cardiology'),
('superadmin', 'super123', 'superadmin', 'System Administrator', 'sysadmin@carehub.com', '555-100-0003', 'IT'),
('nurse', 'nurse123', 'nurse', 'Nancy White', 'nancy.white@carehub.com', '555-100-0004', 'Emergency'),
('intern', 'intern123', 'intern', 'Dr. Michael Lee', 'michael.lee@carehub.com', '555-100-0005', 'Pediatrics'),
('patient', 'patient123', 'patient', 'John Doe', 'john.doe@example.com', '555-123-4567', '');

-- Insert Patients
INSERT INTO patients (first_name, last_name, date_of_birth, email, phone, address) VALUES
('John', 'Doe', '1980-05-15', 'john.doe@example.com', '555-123-4567', '123 Main St, Anytown, CA'),
('Jane', 'Smith', '1975-08-21', 'jane.smith@example.com', '555-987-6543', '456 Oak Ave, Somewhere, NY');

-- Insert Doctors
INSERT INTO doctors (name, role, email, phone, department, specialization, bio, profile_image) VALUES
('Dr. Jane Smith', 'doctor', 'jane.smith@carehub.com', '555-123-4567', 'Cardiology', 'Cardiology', 'Dr. Smith is a board-certified cardiologist with over 15 years of experience.', ''),
('Dr. Robert Chen', 'doctor', 'robert.chen@carehub.com', '555-234-5678', 'Neurology', 'Neurology', 'Dr. Chen specializes in neurological disorders.', 'https://randomuser.me/api/portraits/men/32.jpg'),
('Dr. Maria Rodriguez', 'doctor', 'maria.rodriguez@carehub.com', '555-345-6789', 'Pediatrics', 'Pediatrics', 'Dr. Rodriguez has dedicated her career to children''s health.', 'https://randomuser.me/api/portraits/women/45.jpg'),
('Dr. James Wilson', 'doctor', 'james.wilson@carehub.com', '555-456-7890', 'Orthopedics', 'Orthopedic Surgery', 'Dr. Wilson is an orthopedic surgeon specializing in sports injuries.', '');

-- Insert Appointments
INSERT INTO appointments (patient_id, date_time, description, status, doctor) VALUES
(1, NOW() + INTERVAL 2 DAY, 'Annual checkup', 'Scheduled', 'Dr. Brown'),
(2, NOW() + INTERVAL 3 DAY, 'Follow-up', 'Scheduled', 'Dr. Johnson');

-- Insert Health Metrics
INSERT INTO health_metrics (patient_id, type, value, unit) VALUES
(1, 'Blood Pressure', 120.80, 'mmHg'),
(1, 'Heart Rate', 72, 'bpm'),
(2, 'Blood Pressure', 118.75, 'mmHg');

-- Insert Blogs
INSERT INTO blogs (title, content, excerpt, cover_image, author_id, author_name, published_at, tags) VALUES
('Heart Health Tips', 'Eat well, exercise, and manage stress.', 'Stay heart-healthy!', '', 1, 'Dr. Jane Smith', '2024-06-11 11:00:00', '["cardiology"]'),
('Understanding Migraines', 'Migraines affect millions. Here''s what helps.', 'Guide to controlling migraines.', '', 2, 'Dr. Robert Chen', '2024-05-21 09:30:00', '["neurology"]');

-- Insert Interns
INSERT INTO interns (name, email, department) VALUES
('Alex Green', 'alex.green@carehub.com', 'Cardiology'),
('Priya Patel', 'priya.patel@carehub.com', 'Pediatrics');
