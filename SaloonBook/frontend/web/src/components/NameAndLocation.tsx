import { useState } from "react";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { ArrowLeft, User, MapPin } from "lucide-react";

interface NameAndLocationProps {
  onBack: () => void;
  onComplete: (name: string) => void;
}

export default function NameAndLocation({ onBack, onComplete }: NameAndLocationProps) {
  const [name, setName] = useState("");
  const [locationGranted, setLocationGranted] = useState(false);

  const handleRequestLocation = () => {
    console.log('Requesting location permission');
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          console.log('Location granted:', position.coords);
          setLocationGranted(true);
        },
        (error) => {
          console.log('Location denied:', error);
          setLocationGranted(true);
        }
      );
    } else {
      setLocationGranted(true);
    }
  };

  const handleContinue = () => {
    if (name.trim()) {
      onComplete(name);
    }
  };

  return (
    <div className="min-h-screen flex flex-col bg-background">
      <header className="flex items-center p-4 border-b">
        <Button variant="ghost" size="icon" onClick={onBack} data-testid="button-back">
          <ArrowLeft className="h-5 w-5" />
        </Button>
        <h2 className="ml-4 text-lg font-semibold">Complete Your Profile</h2>
      </header>

      <div className="flex-1 flex items-center justify-center p-6">
        <div className="w-full max-w-md space-y-6">
          <Card className="p-6 space-y-6">
            <div className="space-y-2 text-center">
              <div className="h-16 w-16 rounded-full bg-primary/10 flex items-center justify-center mx-auto">
                <User className="h-8 w-8 text-primary" />
              </div>
              <h3 className="text-xl font-semibold">Tell Us Your Name</h3>
              <p className="text-sm text-muted-foreground">We'll use this to personalize your experience</p>
            </div>

            <div className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="name">Full Name</Label>
                <Input
                  id="name"
                  placeholder="Enter your name"
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                  data-testid="input-name"
                />
              </div>
            </div>
          </Card>

          <Card className="p-6 space-y-4">
            <div className="flex items-start gap-4">
              <div className="h-12 w-12 rounded-full bg-primary/10 flex items-center justify-center flex-shrink-0">
                <MapPin className="h-6 w-6 text-primary" />
              </div>
              <div className="flex-1 space-y-2">
                <h4 className="font-semibold">Enable Location</h4>
                <p className="text-sm text-muted-foreground">
                  We'll use your location to find salons near you
                </p>
                {locationGranted ? (
                  <div className="flex items-center gap-2 text-sm text-primary">
                    <MapPin className="h-4 w-4" />
                    <span>Location enabled</span>
                  </div>
                ) : (
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={handleRequestLocation}
                    data-testid="button-enable-location"
                  >
                    Enable Location
                  </Button>
                )}
              </div>
            </div>
          </Card>

          <Button
            className="w-full"
            onClick={handleContinue}
            disabled={!name.trim() || !locationGranted}
            data-testid="button-continue"
          >
            Continue
          </Button>
        </div>
      </div>
    </div>
  );
}
