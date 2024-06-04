import { TableProducts } from "./components/Table";

const Dashboard = () => {
  return (
    <main className="max-w-6xl mx-auto border shadow p-5 rounded-xl">
      <div className="mb-3">
        <h1 className="text-2xl">Products</h1>
        <p className="text-xs text-muted-foreground">
          These are the details of the product you have.
        </p>
      </div>
      <TableProducts />
    </main>
  );
};

export default Dashboard;
