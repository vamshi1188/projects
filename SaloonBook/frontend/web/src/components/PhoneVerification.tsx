import { useState } from "react";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { ArrowLeft, Phone } from "lucide-react";
import { InputOTP, InputOTPGroup, InputOTPSlot } from "@/components/ui/input-otp";

interface PhoneVerificationProps {
  onBack: () => void;
  onVerified: () => void;
}

export default function PhoneVerification({ onBack, onVerified }: PhoneVerificationProps) {
  const [step, setStep] = useState<"phone" | "otp">("phone");
  const [phoneNumber, setPhoneNumber] = useState("");
  const [otp, setOtp] = useState("");

  const handleSendOTP = () => {
    if (phoneNumber.length >= 10) {
      console.log('Sending OTP to:', phoneNumber);
      setStep("otp");
    }
  };

  const handleVerifyOTP = () => {
    if (otp.length === 6) {
      console.log('Verifying OTP:', otp);
      onVerified();
    }
  };

  return (
    <div className="min-h-screen flex flex-col bg-background">
      <header className="flex items-center p-4 border-b">
        <Button variant="ghost" size="icon" onClick={onBack} data-testid="button-back">
          <ArrowLeft className="h-5 w-5" />
        </Button>
        <h2 className="ml-4 text-lg font-semibold">Phone Verification</h2>
      </header>

      <div className="flex-1 flex items-center justify-center p-6">
        <div className="w-full max-w-md space-y-6">
          {step === "phone" ? (
            <Card className="p-6 space-y-6">
              <div className="space-y-2 text-center">
                <div className="h-16 w-16 rounded-full bg-primary/10 flex items-center justify-center mx-auto">
                  <Phone className="h-8 w-8 text-primary" />
                </div>
                <h3 className="text-xl font-semibold">Enter Your Phone Number</h3>
                <p className="text-sm text-muted-foreground">We'll send you a verification code</p>
              </div>

              <div className="space-y-4">
                <div className="space-y-2">
                  <Label htmlFor="phone">Phone Number</Label>
                  <div className="flex gap-2">
                    <Input
                      id="country-code"
                      value="+91"
                      disabled
                      className="w-20"
                      data-testid="input-country-code"
                    />
                    <Input
                      id="phone"
                      type="tel"
                      placeholder="Enter phone number"
                      value={phoneNumber}
                      onChange={(e) => setPhoneNumber(e.target.value)}
                      data-testid="input-phone"
                    />
                  </div>
                </div>

                <Button className="w-full" onClick={handleSendOTP} data-testid="button-send-otp">
                  Send OTP
                </Button>
              </div>
            </Card>
          ) : (
            <Card className="p-6 space-y-6">
              <div className="space-y-2 text-center">
                <h3 className="text-xl font-semibold">Enter Verification Code</h3>
                <p className="text-sm text-muted-foreground">
                  We sent a code to +91 {phoneNumber}
                </p>
              </div>

              <div className="space-y-4">
                <div className="flex justify-center">
                  <InputOTP maxLength={6} value={otp} onChange={setOtp} data-testid="input-otp">
                    <InputOTPGroup>
                      <InputOTPSlot index={0} />
                      <InputOTPSlot index={1} />
                      <InputOTPSlot index={2} />
                      <InputOTPSlot index={3} />
                      <InputOTPSlot index={4} />
                      <InputOTPSlot index={5} />
                    </InputOTPGroup>
                  </InputOTP>
                </div>

                <Button className="w-full" onClick={handleVerifyOTP} disabled={otp.length !== 6} data-testid="button-verify-otp">
                  Verify OTP
                </Button>

                <Button variant="ghost" className="w-full" onClick={() => setStep("phone")} data-testid="button-change-number">
                  Change Phone Number
                </Button>
              </div>
            </Card>
          )}
        </div>
      </div>
    </div>
  );
}
