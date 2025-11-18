import OrderSummary from '../OrderSummary';
import haircutFade from "@assets/generated_images/Classic_fade_haircut_789ad47e.png";
import beardFull from "@assets/generated_images/Full_beard_style_b3eafe16.png";

export default function OrderSummaryExample() {
  const selectedServices = [
    {
      serviceType: "haircut" as const,
      serviceName: "Haircut",
      styles: [
        {
          id: "fade",
          name: "Classic Fade",
          description: "Timeless short sides with textured top",
          price: 500,
          image: haircutFade,
        },
      ],
    },
    {
      serviceType: "beard" as const,
      serviceName: "Beard",
      styles: [
        {
          id: "full",
          name: "Full Beard",
          description: "Complete beard shaping and styling",
          price: 300,
          image: beardFull,
        },
      ],
    },
  ];

  return (
    <OrderSummary
      selectedServices={selectedServices}
      onBack={() => console.log('Back clicked')}
      onRemoveStyle={(serviceType, styleId) => console.log('Remove:', serviceType, styleId)}
      onProceedToPayment={() => console.log('Proceed to payment')}
    />
  );
}
