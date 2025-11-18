import PaymentInterface from '../PaymentInterface';

export default function PaymentInterfaceExample() {
  return (
    <PaymentInterface
      amount={944}
      onBack={() => console.log('Back clicked')}
      onPaymentSuccess={() => console.log('Payment successful')}
    />
  );
}
