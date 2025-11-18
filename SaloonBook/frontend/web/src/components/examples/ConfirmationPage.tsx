import ConfirmationPage from '../ConfirmationPage';

export default function ConfirmationPageExample() {
  return (
    <ConfirmationPage
      bookingId="RNL2025001"
      userName="John Doe"
      services={[
        { name: "Classic Fade Haircut", price: 500 },
        { name: "Full Beard Styling", price: 300 },
      ]}
      totalAmount={944}
      onBackToHome={() => console.log('Back to home')}
    />
  );
}
