
import { useAuth } from "../context/AuthContext";
import { Button } from "@/components/ui/button";
import { useNavigate } from "react-router-dom";
import Layout from "@/components/Layout";
import { Hospital, ChartBar, Calendar, User } from "lucide-react";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "@/components/ui/carousel";

const Index = () => {
  const { currentUser } = useAuth();
  const navigate = useNavigate();

  const handleGetStarted = () => {
    if (currentUser) {
      navigate("/dashboard");
    } else {
      navigate("/login");
    }
  };

  const testimonials = [
    {
      name: "Sarah Johnson",
      role: "Patient",
      content: "The care I received at CareHub was exceptional. The staff was attentive and professional."
    },
    {
      name: "Dr. Michael Chen",
      role: "Cardiologist",
      content: "State-of-the-art facilities and a collaborative environment make CareHub an excellent place to practice medicine."
    },
    {
      name: "Emma Thompson",
      role: "Healthcare Administrator",
      content: "CareHub's integrated systems have revolutionized how we manage patient care and administrative tasks."
    }
  ];

  return (
    <Layout>
      <div className="flex flex-col min-h-screen">
        {/* Hero Section */}
        <div className="relative bg-gradient-to-r from-teal-600 to-teal-700 text-white py-20">
          <div className="container mx-auto px-4">
            <div className="max-w-3xl mx-auto text-center">
              <h1 className="text-4xl sm:text-5xl md:text-6xl font-extrabold tracking-tight mb-6">
                World-Class Healthcare
                <span className="block mt-2">At Your Service</span>
              </h1>
              <p className="text-xl text-teal-100 mb-8">
                Experience exceptional medical care with cutting-edge technology and compassionate professionals.
              </p>
              <Button 
                onClick={handleGetStarted}
                size="lg"
                className="bg-white text-teal-600 hover:bg-teal-50"
              >
                {currentUser ? "Go to Dashboard" : "Book an Appointment"}
              </Button>
            </div>
          </div>
        </div>

        {/* Statistics Section */}
        <div className="py-16 bg-gray-50">
          <div className="container mx-auto px-4">
            <div className="grid grid-cols-1 md:grid-cols-4 gap-6">
              <Card>
                <CardContent className="pt-6">
                  <div className="flex items-center gap-4">
                    <div className="p-3 bg-teal-100 rounded-full">
                      <Hospital className="w-6 h-6 text-teal-600" />
                    </div>
                    <div>
                      <div className="text-2xl font-bold">50+</div>
                      <div className="text-muted-foreground">Departments</div>
                    </div>
                  </div>
                </CardContent>
              </Card>
              
              <Card>
                <CardContent className="pt-6">
                  <div className="flex items-center gap-4">
                    <div className="p-3 bg-teal-100 rounded-full">
                      <User className="w-6 h-6 text-teal-600" />
                    </div>
                    <div>
                      <div className="text-2xl font-bold">200+</div>
                      <div className="text-muted-foreground">Expert Doctors</div>
                    </div>
                  </div>
                </CardContent>
              </Card>
              
              <Card>
                <CardContent className="pt-6">
                  <div className="flex items-center gap-4">
                    <div className="p-3 bg-teal-100 rounded-full">
                      <Calendar className="w-6 h-6 text-teal-600" />
                    </div>
                    <div>
                      <div className="text-2xl font-bold">10k+</div>
                      <div className="text-muted-foreground">Annual Patients</div>
                    </div>
                  </div>
                </CardContent>
              </Card>
              
              <Card>
                <CardContent className="pt-6">
                  <div className="flex items-center gap-4">
                    <div className="p-3 bg-teal-100 rounded-full">
                      <ChartBar className="w-6 h-6 text-teal-600" />
                    </div>
                    <div>
                      <div className="text-2xl font-bold">95%</div>
                      <div className="text-muted-foreground">Success Rate</div>
                    </div>
                  </div>
                </CardContent>
              </Card>
            </div>
          </div>
        </div>

        {/* Services Section */}
        <div className="py-16">
          <div className="container mx-auto px-4">
            <div className="text-center mb-12">
              <h2 className="text-3xl font-bold mb-4">Our Services</h2>
              <p className="text-muted-foreground max-w-2xl mx-auto">
                Comprehensive healthcare services delivered by experienced professionals using state-of-the-art technology.
              </p>
            </div>
            
            <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
              <Card>
                <CardHeader>
                  <CardTitle>Primary Care</CardTitle>
                  <CardDescription>
                    Comprehensive health assessments and preventive care for all ages
                  </CardDescription>
                </CardHeader>
                <CardContent>
                  <ul className="list-disc list-inside space-y-2 text-muted-foreground">
                    <li>Regular check-ups</li>
                    <li>Preventive screenings</li>
                    <li>Immunizations</li>
                    <li>Health counseling</li>
                  </ul>
                </CardContent>
              </Card>

              <Card>
                <CardHeader>
                  <CardTitle>Specialized Care</CardTitle>
                  <CardDescription>
                    Expert treatment in various medical specialties
                  </CardDescription>
                </CardHeader>
                <CardContent>
                  <ul className="list-disc list-inside space-y-2 text-muted-foreground">
                    <li>Cardiology</li>
                    <li>Orthopedics</li>
                    <li>Neurology</li>
                    <li>Oncology</li>
                  </ul>
                </CardContent>
              </Card>

              <Card>
                <CardHeader>
                  <CardTitle>Emergency Care</CardTitle>
                  <CardDescription>
                    24/7 emergency medical services
                  </CardDescription>
                </CardHeader>
                <CardContent>
                  <ul className="list-disc list-inside space-y-2 text-muted-foreground">
                    <li>Trauma care</li>
                    <li>Critical care</li>
                    <li>Emergency surgery</li>
                    <li>Rapid response</li>
                  </ul>
                </CardContent>
              </Card>
            </div>
          </div>
        </div>

        {/* Testimonials Section */}
        <div className="py-16 bg-gray-50">
          <div className="container mx-auto px-4">
            <div className="text-center mb-12">
              <h2 className="text-3xl font-bold mb-4">What Our Patients Say</h2>
              <p className="text-muted-foreground max-w-2xl mx-auto">
                Read about the experiences of our patients and healthcare professionals.
              </p>
            </div>

            <Carousel className="max-w-xl mx-auto">
              <CarouselContent>
                {testimonials.map((testimonial, index) => (
                  <CarouselItem key={index}>
                    <Card>
                      <CardContent className="pt-6">
                        <blockquote className="space-y-4">
                          <p className="text-muted-foreground">"{testimonial.content}"</p>
                          <footer>
                            <div className="font-semibold">{testimonial.name}</div>
                            <div className="text-sm text-muted-foreground">{testimonial.role}</div>
                          </footer>
                        </blockquote>
                      </CardContent>
                    </Card>
                  </CarouselItem>
                ))}
              </CarouselContent>
              <CarouselPrevious />
              <CarouselNext />
            </Carousel>
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default Index;
