import { RegisterForm } from "@/components/Form";

const Register = () => {
  return (
    <div className="max-w-2xl mx-auto items-center border shadow p-5 rounded-xl">
      <h1 className="text-2xl mb-5">Create your account</h1>
      <RegisterForm />
    </div>
  );
};

export default Register;
