import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Separator } from "@/components/ui/separator";
import { CheckCircle2, Download, Share2 } from "lucide-react";
import { QRCodeSVG } from "qrcode.react";
import successIcon from "@assets/generated_images/Success_confirmation_icon_eba2c618.png";

interface ConfirmationPageProps {
  bookingId: string;
  userName: string;
  services: Array<{ name: string; price: number }>;
  totalAmount: number;
  onBackToHome: () => void;
}

export default function ConfirmationPage({ bookingId, userName, services, totalAmount, onBackToHome }: ConfirmationPageProps) {
  const handleDownloadQR = () => {
    console.log('Download QR code');
  };

  const handleShare = () => {
    console.log('Share booking');
  };

  return (
    <div className="min-h-screen flex flex-col bg-background">
      <div className="flex-1 p-6 space-y-6">
        <div className="text-center space-y-4 pt-8">
          <img src={successIcon} alt="Success" className="h-24 w-24 mx-auto" />
          <div className="space-y-2">
            <h1 className="text-2xl font-bold text-primary">Booking Confirmed!</h1>
            <p className="text-muted-foreground">Your appointment has been successfully booked</p>
          </div>
        </div>

        <Card className="p-6 space-y-4">
          <div className="flex items-center justify-between">
            <h3 className="font-semibold">Booking ID</h3>
            <Badge variant="secondary" className="text-base" data-testid="text-booking-id">{bookingId}</Badge>
          </div>
          <Separator />
          <div className="space-y-2">
            <div className="flex justify-between text-sm">
              <span className="text-muted-foreground">Name</span>
              <span className="font-medium">{userName}</span>
            </div>
            <div className="flex justify-between text-sm">
              <span className="text-muted-foreground">Services</span>
              <span className="font-medium">{services.length} items</span>
            </div>
            <div className="flex justify-between text-sm">
              <span className="text-muted-foreground">Total Amount</span>
              <span className="font-bold text-primary">₹{totalAmount}</span>
            </div>
          </div>
        </Card>

        <Card className="p-6 space-y-4">
          <h3 className="font-semibold text-center">Your Booking QR Code</h3>
          <div className="flex justify-center p-4 bg-white rounded-lg">
            <QRCodeSVG
              value={`RNL-BOOKING-${bookingId}`}
              size={200}
              level="H"
              data-testid="qr-code"
            />
          </div>
          <p className="text-sm text-muted-foreground text-center">
            Show this QR code at the salon
          </p>
          <div className="flex gap-2">
            <Button variant="outline" className="flex-1" onClick={handleDownloadQR} data-testid="button-download-qr">
              <Download className="h-4 w-4 mr-2" />
              Download
            </Button>
            <Button variant="outline" className="flex-1" onClick={handleShare} data-testid="button-share">
              <Share2 className="h-4 w-4 mr-2" />
              Share
            </Button>
          </div>
        </Card>

        <Card className="p-4 space-y-3">
          <h4 className="font-semibold text-sm">Booked Services</h4>
          {services.map((service, index) => (
            <div key={index} className="flex justify-between text-sm">
              <span className="text-muted-foreground">{service.name}</span>
              <span>₹{service.price}</span>
            </div>
          ))}
        </Card>

        <div className="flex items-start gap-3 p-4 bg-primary/5 rounded-lg">
          <CheckCircle2 className="h-5 w-5 text-primary flex-shrink-0 mt-0.5" />
          <div className="flex-1 text-sm">
            <p className="font-medium">Next Steps</p>
            <p className="text-muted-foreground mt-1">
              Visit the salon at your convenience and show this QR code to begin your service.
            </p>
          </div>
        </div>
      </div>

      <div className="sticky bottom-0 p-4 bg-background border-t">
        <Button className="w-full" onClick={onBackToHome} data-testid="button-back-home">
          Back to Home
        </Button>
      </div>
    </div>
  );
}
