import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Separator } from "@/components/ui/separator";
import { ArrowLeft, Trash2 } from "lucide-react";
import type { Style } from "./ServiceSelection";
import type { ServiceType } from "./HomePage";

export interface SelectedService {
  serviceType: ServiceType;
  serviceName: string;
  styles: Style[];
}

interface OrderSummaryProps {
  selectedServices: SelectedService[];
  onBack: () => void;
  onRemoveStyle: (serviceType: ServiceType, styleId: string) => void;
  onProceedToPayment: () => void;
}

export default function OrderSummary({ selectedServices, onBack, onRemoveStyle, onProceedToPayment }: OrderSummaryProps) {
  const subtotal = selectedServices.reduce(
    (sum, service) => sum + service.styles.reduce((s, style) => s + style.price, 0),
    0
  );
  const tax = Math.round(subtotal * 0.18);
  const total = subtotal + tax;

  const totalItems = selectedServices.reduce((sum, service) => sum + service.styles.length, 0);

  return (
    <div className="min-h-screen flex flex-col bg-background">
      <header className="flex items-center p-4 border-b">
        <Button variant="ghost" size="icon" onClick={onBack} data-testid="button-back">
          <ArrowLeft className="h-5 w-5" />
        </Button>
        <h2 className="ml-4 text-lg font-semibold">Order Summary</h2>
      </header>

      <div className="flex-1 p-6 space-y-6">
        <div className="space-y-4">
          {selectedServices.map((service) =>
            service.styles.map((style) => (
              <Card key={`${service.serviceType}-${style.id}`} className="p-4" data-testid={`card-order-item-${style.id}`}>
                <div className="flex gap-4">
                  <img
                    src={style.image}
                    alt={style.name}
                    className="w-20 h-20 rounded-md object-cover"
                  />
                  <div className="flex-1 space-y-1">
                    <div className="flex items-start justify-between gap-2">
                      <div>
                        <Badge variant="outline" className="mb-1">
                          {service.serviceName}
                        </Badge>
                        <h3 className="font-semibold">{style.name}</h3>
                        <p className="text-sm text-muted-foreground">{style.description}</p>
                      </div>
                      <Button
                        variant="ghost"
                        size="icon"
                        onClick={() => onRemoveStyle(service.serviceType, style.id)}
                        data-testid={`button-remove-${style.id}`}
                      >
                        <Trash2 className="h-4 w-4 text-destructive" />
                      </Button>
                    </div>
                    <p className="text-lg font-semibold">₹{style.price}</p>
                  </div>
                </div>
              </Card>
            ))
          )}
        </div>

        <Card className="p-6 space-y-4">
          <h3 className="font-semibold text-lg">Price Details</h3>
          <Separator />
          <div className="space-y-2">
            <div className="flex justify-between">
              <span className="text-muted-foreground">Subtotal ({totalItems} items)</span>
              <span>₹{subtotal}</span>
            </div>
            <div className="flex justify-between">
              <span className="text-muted-foreground">GST (18%)</span>
              <span>₹{tax}</span>
            </div>
          </div>
          <Separator />
          <div className="flex justify-between items-center">
            <span className="text-xl font-bold">Total</span>
            <span className="text-2xl font-bold text-primary" data-testid="text-total-price">₹{total}</span>
          </div>
        </Card>
      </div>

      <div className="sticky bottom-0 p-4 bg-background border-t">
        <Button className="w-full" onClick={onProceedToPayment} data-testid="button-proceed-payment">
          Proceed to Payment
        </Button>
      </div>
    </div>
  );
}
