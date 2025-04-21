
import { useState, useEffect } from "react";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Button } from "@/components/ui/button";
import { patientService } from "../../services/api";
import { useToast } from "@/components/ui/use-toast";
import PatientForm from "./PatientForm";
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog";

interface Patient {
  id: number;
  firstName: string;
  lastName: string;
  dateOfBirth: string;
  email: string;
  phone: string;
  address: string;
  createdAt: string;
}

export default function PatientList() {
  const [patients, setPatients] = useState<Patient[]>([]);
  const [loading, setLoading] = useState(true);
  const [openDialog, setOpenDialog] = useState(false);
  const [currentPatient, setCurrentPatient] = useState<Patient | null>(null);
  const { toast } = useToast();

  const fetchPatients = async () => {
    try {
      setLoading(true);
      const data = await patientService.getAllPatients();
      setPatients(data);
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to fetch patients",
        variant: "destructive",
      });
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchPatients();
  }, []);

  const handleAddEdit = (patient: Patient | null) => {
    setCurrentPatient(patient);
    setOpenDialog(true);
  };

  const handleSave = async (patientData: any) => {
    try {
      if (currentPatient) {
        await patientService.updatePatient(currentPatient.id, patientData);
        toast({
          title: "Success",
          description: "Patient updated successfully",
        });
      } else {
        await patientService.createPatient(patientData);
        toast({
          title: "Success",
          description: "Patient created successfully",
        });
      }
      setOpenDialog(false);
      fetchPatients();
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to save patient",
        variant: "destructive",
      });
    }
  };

  const handleDelete = async (id: number) => {
    if (confirm("Are you sure you want to delete this patient?")) {
      try {
        await patientService.deletePatient(id);
        toast({
          title: "Success",
          description: "Patient deleted successfully",
        });
        fetchPatients();
      } catch (error) {
        toast({
          title: "Error",
          description: "Failed to delete patient",
          variant: "destructive",
        });
      }
    }
  };

  if (loading) {
    return <div className="flex justify-center p-4">Loading patients...</div>;
  }

  return (
    <>
      <div className="flex justify-between items-center mb-4">
        <h2 className="text-xl font-semibold">Patient Records</h2>
        <Button 
          onClick={() => handleAddEdit(null)}
          className="bg-teal-600 hover:bg-teal-700"
        >
          Add New Patient
        </Button>
      </div>

      <div className="border rounded-md">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead>Date of Birth</TableHead>
              <TableHead>Contact</TableHead>
              <TableHead className="hidden md:table-cell">Address</TableHead>
              <TableHead>Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {patients.length === 0 ? (
              <TableRow>
                <TableCell colSpan={5} className="text-center">No patients found</TableCell>
              </TableRow>
            ) : (
              patients.map((patient) => (
                <TableRow key={patient.id}>
                  <TableCell>
                    <div className="font-medium">{patient.firstName} {patient.lastName}</div>
                  </TableCell>
                  <TableCell>{patient.dateOfBirth}</TableCell>
                  <TableCell>
                    <div>{patient.email}</div>
                    <div className="text-sm text-muted-foreground">{patient.phone}</div>
                  </TableCell>
                  <TableCell className="hidden md:table-cell">{patient.address}</TableCell>
                  <TableCell>
                    <div className="flex space-x-2">
                      <Button 
                        variant="outline" 
                        size="sm"
                        onClick={() => handleAddEdit(patient)}
                      >
                        Edit
                      </Button>
                      <Button 
                        variant="destructive" 
                        size="sm"
                        onClick={() => handleDelete(patient.id)}
                      >
                        Delete
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
            <DialogTitle>{currentPatient ? "Edit Patient" : "Add New Patient"}</DialogTitle>
          </DialogHeader>
          <PatientForm 
            patient={currentPatient} 
            onSave={handleSave} 
            onCancel={() => setOpenDialog(false)} 
          />
        </DialogContent>
      </Dialog>
    </>
  );
}
