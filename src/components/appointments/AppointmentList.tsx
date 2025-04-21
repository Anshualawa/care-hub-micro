
import { useState, useEffect } from "react";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Button } from "@/components/ui/button";
import { appointmentService, patientService } from "../../services/api";
import { useToast } from "@/components/ui/use-toast";
import AppointmentForm from "./AppointmentForm";
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { Badge } from "@/components/ui/badge";
import { formatDate } from "../../utils/dateUtils";

interface Appointment {
  id: number;
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

export default function AppointmentList() {
  const [appointments, setAppointments] = useState<Appointment[]>([]);
  const [patients, setPatients] = useState<Patient[]>([]);
  const [loading, setLoading] = useState(true);
  const [openDialog, setOpenDialog] = useState(false);
  const [currentAppointment, setCurrentAppointment] = useState<Appointment | null>(null);
  const { toast } = useToast();

  const fetchData = async () => {
    try {
      setLoading(true);
      const [appointmentsData, patientsData] = await Promise.all([
        appointmentService.getAllAppointments(),
        patientService.getAllPatients()
      ]);
      
      setAppointments(appointmentsData);
      setPatients(patientsData);
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to fetch data",
        variant: "destructive",
      });
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  const handleAddEdit = (appointment: Appointment | null) => {
    setCurrentAppointment(appointment);
    setOpenDialog(true);
  };

  const handleSave = async (appointmentData: any) => {
    try {
      if (currentAppointment) {
        await appointmentService.updateAppointment(currentAppointment.id, appointmentData);
        toast({
          title: "Success",
          description: "Appointment updated successfully",
        });
      } else {
        await appointmentService.createAppointment(appointmentData);
        toast({
          title: "Success",
          description: "Appointment created successfully",
        });
      }
      setOpenDialog(false);
      fetchData();
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to save appointment",
        variant: "destructive",
      });
    }
  };

  const handleDelete = async (id: number) => {
    if (confirm("Are you sure you want to delete this appointment?")) {
      try {
        await appointmentService.deleteAppointment(id);
        toast({
          title: "Success",
          description: "Appointment deleted successfully",
        });
        fetchData();
      } catch (error) {
        toast({
          title: "Error",
          description: "Failed to delete appointment",
          variant: "destructive",
        });
      }
    }
  };

  const getPatientName = (patientId: number) => {
    const patient = patients.find(p => p.id === patientId);
    return patient ? `${patient.firstName} ${patient.lastName}` : "Unknown Patient";
  };

  const getStatusColor = (status: string) => {
    switch (status.toLowerCase()) {
      case "scheduled":
        return "bg-blue-500";
      case "completed":
        return "bg-green-500";
      case "cancelled":
        return "bg-red-500";
      default:
        return "bg-gray-500";
    }
  };

  if (loading) {
    return <div className="flex justify-center p-4">Loading appointments...</div>;
  }

  return (
    <>
      <div className="flex justify-between items-center mb-4">
        <h2 className="text-xl font-semibold">Appointment Schedule</h2>
        <Button 
          onClick={() => handleAddEdit(null)}
          className="bg-teal-600 hover:bg-teal-700"
        >
          Schedule Appointment
        </Button>
      </div>

      <div className="border rounded-md">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Patient</TableHead>
              <TableHead>Date & Time</TableHead>
              <TableHead>Doctor</TableHead>
              <TableHead>Status</TableHead>
              <TableHead>Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {appointments.length === 0 ? (
              <TableRow>
                <TableCell colSpan={5} className="text-center">No appointments found</TableCell>
              </TableRow>
            ) : (
              appointments.map((appointment) => (
                <TableRow key={appointment.id}>
                  <TableCell className="font-medium">
                    {getPatientName(appointment.patientId)}
                  </TableCell>
                  <TableCell>
                    <div className="font-medium">{new Date(appointment.dateTime).toLocaleDateString()}</div>
                    <div className="text-sm text-muted-foreground">{new Date(appointment.dateTime).toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'})}</div>
                  </TableCell>
                  <TableCell>{appointment.doctor}</TableCell>
                  <TableCell>
                    <Badge className={getStatusColor(appointment.status)}>
                      {appointment.status}
                    </Badge>
                  </TableCell>
                  <TableCell>
                    <div className="flex space-x-2">
                      <Button 
                        variant="outline" 
                        size="sm"
                        onClick={() => handleAddEdit(appointment)}
                      >
                        Edit
                      </Button>
                      <Button 
                        variant="destructive" 
                        size="sm"
                        onClick={() => handleDelete(appointment.id)}
                      >
                        Cancel
                      </Button>
                    </div>
                  </TableCell>
                </TableRow>
              ))
            )}
          </TableBody>
        </Table>
      </div>

      <Dialog open={openDialog} onOpenChange={setOpenDialog}>
        <DialogContent className="sm:max-w-[500px]">
          <DialogHeader>
            <DialogTitle>{currentAppointment ? "Edit Appointment" : "Schedule New Appointment"}</DialogTitle>
          </DialogHeader>
          <AppointmentForm 
            appointment={currentAppointment} 
            patients={patients}
            onSave={handleSave} 
            onCancel={() => setOpenDialog(false)} 
          />
        </DialogContent>
      </Dialog>
    </>
  );
}
