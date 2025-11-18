import { useState } from "react";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { ArrowLeft, CreditCard, Wallet, Building2 } from "lucide-react";

interface PaymentInterfaceProps {
  amount: number;
  onBack: () => void;
  onPaymentSuccess: () => void;
}

export default function PaymentInterface({ amount, onBack, onPaymentSuccess }: PaymentInterfaceProps) {
  const [selectedMethod, setSelectedMethod] = useState<string | null>(null);
  const [processing, setProcessing] = useState(false);

  const paymentMethods = [
    { id: "card", name: "Credit/Debit Card", icon: CreditCard, description: "Visa, Mastercard, RuPay" },
    { id: "upi", name: "UPI", icon: Wallet, description: "Google Pay, PhonePe, Paytm" },
    { id: "netbanking", name: "Net Banking", icon: Building2, description: "All major banks" },
  ];

  const handlePayment = () => {
    if (!selectedMethod) return;
    
    setProcessing(true);
    console.log('Processing payment with:', selectedMethod);
    
    setTimeout(() => {
      setProcessing(false);
      onPaymentSuccess();
    }, 2000);
  };

  return (
    <div className="min-h-screen flex flex-col bg-background">
      <header className="flex items-center p-4 border-b">
        <Button variant="ghost" size="icon" onClick={onBack} disabled={processing} data-testid="button-back">
          <ArrowLeft className="h-5 w-5" />
        </Button>
        <h2 className="ml-4 text-lg font-semibold">Payment</h2>
      </header>

      <div className="flex-1 p-6 space-y-6">
        <Card className="p-6">
          <div className="text-center space-y-2">
            <p className="text-sm text-muted-foreground">Amount to Pay</p>
            <p className="text-3xl font-bold text-primary" data-testid="text-payment-amount">₹{amount}</p>
          </div>
        </Card>

        <div className="space-y-3">
          <h3 className="font-semibold">Select Payment Method</h3>
          {paymentMethods.map((method) => (
            <Card
              key={method.id}
              className={`p-4 hover-elevate active-elevate-2 cursor-pointer ${
                selectedMethod === method.id ? "ring-2 ring-primary" : ""
              }`}
              onClick={() => setSelectedMethod(method.id)}
              data-testid={`card-payment-${method.id}`}
            >
              <div className="flex items-center gap-4">
                <div className="h-12 w-12 rounded-full bg-primary/10 flex items-center justify-center">
                  <method.icon className="h-6 w-6 text-primary" />
                </div>
                <div className="flex-1">
                  <h4 className="font-semibold">{method.name}</h4>
                  <p className="text-sm text-muted-foreground">{method.description}</p>
                </div>
              </div>
            </Card>
          ))}
        </div>

        <Card className="p-4 bg-muted/50">
          <div className="flex items-center gap-2 text-sm">
            <CreditCard className="h-4 w-4 text-muted-foreground" />
            <span className="text-muted-foreground">Secured by Stripe</span>
          </div>
        </Card>
      </div>

      <div className="sticky bottom-0 p-4 bg-background border-t">
        <Button
          className="w-full"
          onClick={handlePayment}
          disabled={!selectedMethod || processing}
          data-testid="button-pay"
        >
          {processing ? "Processing..." : `Pay ₹${amount}`}
        </Button>
      </div>
    </div>
  );
}
