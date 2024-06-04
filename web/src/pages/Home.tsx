import { Button } from "@/components/ui/button";
import { Link } from "react-router-dom";

const Home = () => {
  return (
    <div className="flex flex-col justify-center items-center">
      <h1>Welcome to my APPS</h1>
      <Link to="/auth/register">
        <Button variant="default">go to register</Button>
      </Link>
    </div>
  );
};

export default Home;
