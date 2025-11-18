import { useState } from "react";
import { Switch, Route } from "wouter";
import { queryClient } from "./lib/queryClient";
import { QueryClientProvider } from "@tanstack/react-query";
import { Toaster } from "@/components/ui/toaster";
import { TooltipProvider } from "@/components/ui/tooltip";
import WelcomeScreen from "@/components/WelcomeScreen";
import PhoneVerification from "@/components/PhoneVerification";
import NameAndLocation from "@/components/NameAndLocation";
import HomePage from "@/components/HomePage";
import ServiceSelection from "@/components/ServiceSelection";
import OrderSummary, { type SelectedService } from "@/components/OrderSummary";
import PaymentInterface from "@/components/PaymentInterface";
import ConfirmationPage from "@/components/ConfirmationPage";
import type { ServiceType } from "@/components/HomePage";
import type { Style } from "@/components/ServiceSelection";
import NotFound from "@/pages/not-found";

type AppScreen =
  | "welcome"
  | "signup-phone"
  | "signup-name"
  | "login"
  | "home"
  | "service-selection"
  | "order-summary"
  | "payment"
  | "confirmation";

const serviceNames: Record<ServiceType, string> = {
  haircut: "Haircut",
  beard: "Beard",
  color: "Hair & Beard Color",
  massage: "Massage",
  facewash: "Face Wash",
};

function BookingApp() {
  const [currentScreen, setCurrentScreen] = useState<AppScreen>("welcome");
  const [userName, setUserName] = useState("");
  const [currentServiceType, setCurrentServiceType] = useState<ServiceType | null>(null);
  const [selectedServices, setSelectedServices] = useState<SelectedService[]>([]);
  const [bookingId] = useState(`RNL${Date.now()}`);

  const serviceQueue: ServiceType[] = ["haircut", "beard", "color", "massage", "facewash"];
  const [currentServiceIndex, setCurrentServiceIndex] = useState(0);

  const handleSignUp = () => {
    setCurrentScreen("signup-phone");
  };

  const handleLogin = () => {
    setUserName("Demo User");
    setCurrentScreen("home");
  };

  const handlePhoneVerified = () => {
    setCurrentScreen("signup-name");
  };

  const handleNameComplete = (name: string) => {
    setUserName(name);
    setCurrentScreen("home");
  };

  const handleServiceSelect = (serviceType: ServiceType) => {
    const index = serviceQueue.indexOf(serviceType);
    setCurrentServiceIndex(index);
    setCurrentServiceType(serviceType);
    setCurrentScreen("service-selection");
  };

  const handleServiceContinue = (styles: Style[]) => {
    if (!currentServiceType) return;

    if (styles.length > 0) {
      const existingServiceIndex = selectedServices.findIndex(
        (s) => s.serviceType === currentServiceType
      );

      if (existingServiceIndex >= 0) {
        const updated = [...selectedServices];
        updated[existingServiceIndex] = {
          ...updated[existingServiceIndex],
          styles: [...updated[existingServiceIndex].styles, ...styles],
        };
        setSelectedServices(updated);
      } else {
        setSelectedServices([
          ...selectedServices,
          {
            serviceType: currentServiceType,
            serviceName: serviceNames[currentServiceType],
            styles,
          },
        ]);
      }
    }

    const nextIndex = currentServiceIndex + 1;
    if (nextIndex < serviceQueue.length) {
      setCurrentServiceIndex(nextIndex);
      setCurrentServiceType(serviceQueue[nextIndex]);
    } else {
      setCurrentScreen("order-summary");
    }
  };

  const handleServiceSkip = () => {
    const nextIndex = currentServiceIndex + 1;
    if (nextIndex < serviceQueue.length) {
      setCurrentServiceIndex(nextIndex);
      setCurrentServiceType(serviceQueue[nextIndex]);
    } else {
      setCurrentScreen("order-summary");
    }
  };

  const handleRemoveStyle = (serviceType: ServiceType, styleId: string) => {
    const updated = selectedServices.map((service) => {
      if (service.serviceType === serviceType) {
        return {
          ...service,
          styles: service.styles.filter((s) => s.id !== styleId),
        };
      }
      return service;
    }).filter((service) => service.styles.length > 0);

    setSelectedServices(updated);
  };

  const handleProceedToPayment = () => {
    setCurrentScreen("payment");
  };

  const handlePaymentSuccess = () => {
    setCurrentScreen("confirmation");
  };

  const handleBackToHome = () => {
    setSelectedServices([]);
    setCurrentServiceIndex(0);
    setCurrentServiceType(null);
    setCurrentScreen("home");
  };

  const totalAmount = selectedServices.reduce(
    (sum, service) => sum + service.styles.reduce((s, style) => s + style.price, 0),
    0
  ) * 1.18;

  const totalAmountRounded = Math.round(totalAmount);

  const flatServiceList = selectedServices.flatMap((service) =>
    service.styles.map((style) => ({
      name: `${style.name} (${service.serviceName})`,
      price: style.price,
    }))
  );

  return (
    <>
      {currentScreen === "welcome" && (
        <WelcomeScreen onSignUp={handleSignUp} onLogin={handleLogin} />
      )}
      {currentScreen === "signup-phone" && (
        <PhoneVerification
          onBack={() => setCurrentScreen("welcome")}
          onVerified={handlePhoneVerified}
        />
      )}
      {currentScreen === "signup-name" && (
        <NameAndLocation
          onBack={() => setCurrentScreen("signup-phone")}
          onComplete={handleNameComplete}
        />
      )}
      {currentScreen === "home" && (
        <HomePage onServiceSelect={handleServiceSelect} userName={userName} />
      )}
      {currentScreen === "service-selection" && currentServiceType && (
        <ServiceSelection
          serviceType={currentServiceType}
          onBack={handleBackToHome}
          onSkip={handleServiceSkip}
          onContinue={handleServiceContinue}
        />
      )}
      {currentScreen === "order-summary" && (
        <OrderSummary
          selectedServices={selectedServices}
          onBack={handleBackToHome}
          onRemoveStyle={handleRemoveStyle}
          onProceedToPayment={handleProceedToPayment}
        />
      )}
      {currentScreen === "payment" && (
        <PaymentInterface
          amount={totalAmountRounded}
          onBack={() => setCurrentScreen("order-summary")}
          onPaymentSuccess={handlePaymentSuccess}
        />
      )}
      {currentScreen === "confirmation" && (
        <ConfirmationPage
          bookingId={bookingId}
          userName={userName}
          services={flatServiceList}
          totalAmount={totalAmountRounded}
          onBackToHome={handleBackToHome}
        />
      )}
    </>
  );
}

function Router() {
  return (
    <Switch>
      <Route path="/" component={BookingApp} />
      <Route component={NotFound} />
    </Switch>
  );
}

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <TooltipProvider>
        <Toaster />
        <Router />
      </TooltipProvider>
    </QueryClientProvider>
  );
}

export default App;
