
import { useState, useEffect } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { healthMetricService } from "../../services/api";
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from "recharts";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Button } from "@/components/ui/button";
import { useToast } from "@/components/ui/use-toast";

interface HealthMetric {
  id: number;
  patientId: number;
  type: string;
  value: number;
  unit: string;
  recordedAt: string;
}

interface HealthMetricsChartProps {
  patientId: number;
}

export default function HealthMetricsChart({ patientId }: HealthMetricsChartProps) {
  const [metrics, setMetrics] = useState<HealthMetric[]>([]);
  const [loading, setLoading] = useState(true);
  const [metricType, setMetricType] = useState("Blood Pressure");
  const { toast } = useToast();

  useEffect(() => {
    fetchMetrics();
  }, [patientId]);

  const fetchMetrics = async () => {
    try {
      setLoading(true);
      const data = await healthMetricService.getPatientMetrics(patientId);
      setMetrics(data);
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to fetch health metrics",
        variant: "destructive",
      });
    } finally {
      setLoading(false);
    }
  };

  const formatChartData = () => {
    const filteredMetrics = metrics.filter(metric => metric.type === metricType);
    
    return filteredMetrics.map(metric => ({
      date: new Date(metric.recordedAt).toLocaleDateString(),
      value: metric.value,
    }));
  };

  const getUnitForCurrentMetric = () => {
    const metric = metrics.find(m => m.type === metricType);
    return metric ? metric.unit : "";
  };

  if (loading) {
    return <div className="flex justify-center p-4">Loading health metrics...</div>;
  }

  if (metrics.length === 0) {
    return (
      <Card className="w-full h-64 flex items-center justify-center">
        <p className="text-muted-foreground">No health metrics available for this patient</p>
      </Card>
    );
  }

  // Get unique metric types
  const metricTypes = Array.from(new Set(metrics.map(metric => metric.type)));

  return (
    <Card className="w-full">
      <CardHeader className="flex flex-row items-center justify-between">
        <CardTitle>Health Metrics</CardTitle>
        <div className="flex space-x-2">
          <Select value={metricType} onValueChange={setMetricType}>
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder="Metric type" />
            </SelectTrigger>
            <SelectContent>
              {metricTypes.map(type => (
                <SelectItem key={type} value={type}>{type}</SelectItem>
              ))}
            </SelectContent>
          </Select>
          <Button variant="outline" onClick={fetchMetrics}>Refresh</Button>
        </div>
      </CardHeader>
      <CardContent>
        <ResponsiveContainer width="100%" height={300}>
          <LineChart
            data={formatChartData()}
            margin={{
              top: 5,
              right: 30,
              left: 20,
              bottom: 5,
            }}
          >
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey="date" />
            <YAxis unit={` ${getUnitForCurrentMetric()}`} />
            <Tooltip />
            <Legend />
            <Line 
              type="monotone" 
              dataKey="value" 
              stroke="#0d9488" 
              activeDot={{ r: 8 }} 
              name={metricType} 
            />
          </LineChart>
        </ResponsiveContainer>
      </CardContent>
    </Card>
  );
}
