import PhoneVerification from '../PhoneVerification';

export default function PhoneVerificationExample() {
  return (
    <PhoneVerification
      onBack={() => console.log('Back clicked')}
      onVerified={() => console.log('Phone verified')}
    />
  );
}
