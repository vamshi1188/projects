import NameAndLocation from '../NameAndLocation';

export default function NameAndLocationExample() {
  return (
    <NameAndLocation
      onBack={() => console.log('Back clicked')}
      onComplete={(name) => console.log('Completed with name:', name)}
    />
  );
}
