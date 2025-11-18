import HomePage from '../HomePage';

export default function HomePageExample() {
  return (
    <HomePage
      onServiceSelect={(service) => console.log('Service selected:', service)}
      userName="John"
    />
  );
}
