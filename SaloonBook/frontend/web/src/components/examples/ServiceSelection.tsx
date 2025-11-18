import ServiceSelection from '../ServiceSelection';

export default function ServiceSelectionExample() {
  return (
    <ServiceSelection
      serviceType="haircut"
      onBack={() => console.log('Back clicked')}
      onSkip={() => console.log('Skip clicked')}
      onContinue={(styles) => console.log('Continue with styles:', styles)}
    />
  );
}
