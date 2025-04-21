
import { useState, useEffect } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";

interface Appointment {
  id?: number;
  patientId: number;
  dateTime: string;
  description: string;
  status: string;
  doctor: string;
}

interface Patient {
  id: number;
  firstName: string;
  lastName: string;
}

interface AppointmentFormProps {
  appointment: Appointment | null;
  patients: Patient[];
  onSave: (appointment: Appointment) => void;
  onCancel: () => void;
}

export default function AppointmentForm({ appointment, patients, onSave, onCancel }: AppointmentFormProps) {
  const [formData, setFormData] = useState<Appointment>({
    patientId: 0,
    dateTime: new Date().toISOString().split('T')[0] + 'T09:00',
    description: "",
    status: "Scheduled",
    doctor: ""
  });

  useEffect(() => {
    if (appointment) {
      // Format the date for the datetime-local input
      const dateTime = new Date(appointment.dateTime)
        .toISOString()
        .slice(0, 16); // Get YYYY-MM-DDTHH:MM format
      
      setFormData({
        ...appointment,
        dateTime
      });
    } else if (patients.length > 0) {
      // Set the first patient as default for new appointments
      setFormData(prev => ({
        ...prev,
        patientId: patients[0].id
      }));
    }
  }, [appointment, patients]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSelectChange = (name: string, value: string) => {
    setFormData({ ...formData, [name]: name === "patientId" ? parseInt(value) : value });
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSave(formData);
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <div className="space-y-2">
        <Label htmlFor="patientId">Patient</Label>
        <Select 
          value={formData.patientId.toString()} 
          onValueChange={(value) => handleSelectChange("patientId", value)}
        >
          <SelectTrigger id="patientId">
            <SelectValue placeholder="Select a patient" />
          </SelectTrigger>
          <SelectContent>
            {patients.map((patient) => (
              <SelectItem key={patient.id} value={patient.id.toString()}>
                {patient.firstName} {patient.lastName}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>

      <div className="space-y-2">
        <Label htmlFor="dateTime">Date & Time</Label>
        <Input
          id="dateTime"
          name="dateTime"
          type="datetime-local"
          value={formData.dateTime}
          onChange={handleChange}
          required
        />
      </div>

      <div className="space-y-2">
        <Label htmlFor="doctor">Doctor</Label>
        <Select 
          value={formData.doctor}
          onValueChange={(value) => handleSelectChange("doctor", value)}
        >
          <SelectTrigger id="doctor">
            <SelectValue placeholder="Select a doctor" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="Dr. Brown">Dr. Brown</SelectItem>
            <SelectItem value="Dr. Johnson">Dr. Johnson</SelectItem>
            <SelectItem value="Dr. Williams">Dr. Williams</SelectItem>
            <SelectItem value="Dr. Davis">Dr. Davis</SelectItem>
          </SelectContent>
        </Select>
      </div>

      <div className="space-y-2">
        <Label htmlFor="status">Status</Label>
        <Select 
          value={formData.status}
          onValueChange={(value) => handleSelectChange("status", value)}
        >
          <SelectTrigger id="status">
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="Scheduled">Scheduled</SelectItem>
            <SelectItem value="Completed">Completed</SelectItem>
            <SelectItem value="Cancelled">Cancelled</SelectItem>
            <SelectItem value="Rescheduled">Rescheduled</SelectItem>
          </SelectContent>
        </Select>
      </div>

      <div className="space-y-2">
        <Label htmlFor="description">Description</Label>
        <Textarea
          id="description"
          name="description"
          value={formData.description}
          onChange={handleChange}
          rows={3}
        />
      </div>

      <div className="flex justify-end space-x-4 pt-4">
        <Button type="button" variant="outline" onClick={onCancel}>
          Cancel
        </Button>
        <Button type="submit" className="bg-teal-600 hover:bg-teal-700">
          {appointment ? "Update Appointment" : "Schedule Appointment"}
        </Button>
      </div>
    </form>
  );
}
