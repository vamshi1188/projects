import { Card } from "@/components/ui/card";
import haircutIcon from "@assets/generated_images/Haircut_service_icon_afbde6f1.png";
import beardIcon from "@assets/generated_images/Beard_service_icon_3acdc600.png";
import colorIcon from "@assets/generated_images/Hair_color_service_icon_39ba00ce.png";
import massageIcon from "@assets/generated_images/Massage_service_icon_97dfc7a3.png";
import faceWashIcon from "@assets/generated_images/Face_wash_service_icon_f8f9a509.png";

export type ServiceType = "haircut" | "beard" | "color" | "massage" | "facewash";

interface HomePageProps {
  onServiceSelect: (service: ServiceType) => void;
  userName?: string;
}

const services = [
  {
    id: "haircut" as ServiceType,
    name: "Haircut",
    description: "Professional styling",
    icon: haircutIcon,
  },
  {
    id: "beard" as ServiceType,
    name: "Beard",
    description: "Grooming & styling",
    icon: beardIcon,
  },
  {
    id: "color" as ServiceType,
    name: "Hair & Beard Color",
    description: "Expert coloring",
    icon: colorIcon,
  },
  {
    id: "massage" as ServiceType,
    name: "Massage",
    description: "Relaxation therapy",
    icon: massageIcon,
  },
  {
    id: "facewash" as ServiceType,
    name: "Face Wash",
    description: "Skincare treatment",
    icon: faceWashIcon,
  },
];

export default function HomePage({ onServiceSelect, userName }: HomePageProps) {
  return (
    <div className="min-h-screen bg-background">
      <header className="p-6 space-y-2">
        <h1 className="text-3xl font-bold">
          Welcome{userName ? `, ${userName}` : ''}
        </h1>
        <p className="text-muted-foreground">Choose your salon service</p>
      </header>

      <div className="p-6 grid grid-cols-2 gap-4">
        {services.map((service) => (
          <Card
            key={service.id}
            className="p-6 space-y-4 hover-elevate active-elevate-2 cursor-pointer"
            onClick={() => onServiceSelect(service.id)}
            data-testid={`card-service-${service.id}`}
          >
            <div className="aspect-square flex items-center justify-center">
              <img src={service.icon} alt={service.name} className="h-20 w-20 object-contain" />
            </div>
            <div className="text-center space-y-1">
              <h3 className="font-semibold text-lg">{service.name}</h3>
              <p className="text-sm text-muted-foreground">{service.description}</p>
            </div>
          </Card>
        ))}
      </div>
    </div>
  );
}
