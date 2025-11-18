import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { UserPlus, LogIn } from "lucide-react";
import logoImage from "@assets/generated_images/RNL_app_logo_7e2a7034.png";

interface WelcomeScreenProps {
  onSignUp: () => void;
  onLogin: () => void;
}

export default function WelcomeScreen({ onSignUp, onLogin }: WelcomeScreenProps) {
  return (
    <div className="min-h-screen flex flex-col items-center justify-center p-6 bg-gradient-to-b from-primary/5 to-background">
      <div className="w-full max-w-md space-y-8">
        <div className="text-center space-y-4">
          <img src={logoImage} alt="RNL Logo" className="h-32 mx-auto" />
          <h1 className="text-3xl font-bold">Welcome to RNL</h1>
          <p className="text-muted-foreground">Your premium salon booking experience</p>
        </div>

        <div className="space-y-4">
          <Card className="p-6 hover-elevate active-elevate-2 cursor-pointer" onClick={onSignUp} data-testid="card-signup">
            <div className="flex items-center gap-4">
              <div className="h-12 w-12 rounded-full bg-primary/10 flex items-center justify-center">
                <UserPlus className="h-6 w-6 text-primary" />
              </div>
              <div className="flex-1">
                <h3 className="font-semibold text-lg">Sign Up</h3>
                <p className="text-sm text-muted-foreground">Create a new account</p>
              </div>
            </div>
          </Card>

          <Card className="p-6 hover-elevate active-elevate-2 cursor-pointer" onClick={onLogin} data-testid="card-login">
            <div className="flex items-center gap-4">
              <div className="h-12 w-12 rounded-full bg-primary/10 flex items-center justify-center">
                <LogIn className="h-6 w-6 text-primary" />
              </div>
              <div className="flex-1">
                <h3 className="font-semibold text-lg">Already a User</h3>
                <p className="text-sm text-muted-foreground">Sign in to your account</p>
              </div>
            </div>
          </Card>
        </div>
      </div>
    </div>
  );
}
