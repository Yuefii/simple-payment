import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import {
  Table,
  TableBody,
  TableCell,
  TableFooter,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { useState } from "react";

const invoices = [
  {
    product_id: "INV001",
    name: "Paid",
    price: "$250.00",
  },
  {
    product_id: "INV002",
    name: "Pending",
    price: "$150.00",
  },
  {
    product_id: "INV003",
    name: "Unpaid",
    price: "$350.00",
  },
  {
    product_id: "INV004",
    name: "Paid",
    price: "$450.00",
  },
  {
    product_id: "INV005",
    name: "Paid",
    price: "$550.00",
  },
  {
    product_id: "INV006",
    name: "Pending",
    price: "$200.00",
  },
  {
    product_id: "INV007",
    name: "Unpaid",
    price: "$300.00",
  },
];

export function TableProducts() {
  const [selectedProducts, setSelectedProducts] = useState<string[]>([]);
  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead className="w-[100px]">
            <Checkbox
              checked={selectedProducts.length === invoices.length}
              onCheckedChange={(checked) => {
                if (checked) {
                  setSelectedProducts(invoices.map((item) => item.product_id));
                } else {
                  setSelectedProducts([]);
                }
                console.log(checked);
              }}
            />
          </TableHead>
          <TableHead>Product ID</TableHead>
          <TableHead>Name</TableHead>
          <TableHead>Price</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {invoices.map((item) => (
          <TableRow key={item.product_id}>
            <TableCell>
              <Checkbox
                checked={selectedProducts.includes(item.product_id)}
                onCheckedChange={(checked) => {
                  if (checked) {
                    setSelectedProducts((prev) => [...prev, item.product_id]);
                  } else {
                    setSelectedProducts((prev) =>
                      prev.filter((id) => id !== item.product_id)
                    );
                  }
                  console.log(checked);
                }}
              />
            </TableCell>
            <TableCell className="font-medium">{item.product_id}</TableCell>
            <TableCell className="font-medium">{item.name}</TableCell>
            <TableCell>{item.price}</TableCell>
          </TableRow>
        ))}
      </TableBody>
      <TableFooter>
        <TableRow>
          <TableCell colSpan={4} className="text-right">
            <Button>Add Products</Button>
          </TableCell>
        </TableRow>
      </TableFooter>
    </Table>
  );
}
