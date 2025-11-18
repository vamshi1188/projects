import WelcomeScreen from '../WelcomeScreen';

export default function WelcomeScreenExample() {
  return (
    <WelcomeScreen
      onSignUp={() => console.log('Sign up clicked')}
      onLogin={() => console.log('Login clicked')}
    />
  );
}
