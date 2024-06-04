import { z } from "zod";
import { Button } from "@/components/ui/button";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Input } from "@/components/ui/input";
import { Login } from "@/validation/form_schema";
import * as ui from "@/components/ui/form";
import axios from "axios";
import { useNavigate } from "react-router-dom";

export function LoginForm() {
  const navigate = useNavigate()
  const form = useForm<z.infer<typeof Login>>({
    resolver: zodResolver(Login),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  const onSubmit = async (values: z.infer<typeof Login>) => {
    try {
      const res = await axios.post(
        `http://localhost:8000/auth/login`,
        values
      );
      localStorage.setItem("token", res.data.token);
      navigate("/");
    } catch (error) {
      console.error("Error during registration:", error);
    }
  };

  return (
    <ui.Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-2">
        <ui.FormField
          control={form.control}
          name="username"
          render={({ field }) => (
            <ui.FormItem>
              <ui.FormLabel>Username</ui.FormLabel>
              <ui.FormControl>
                <Input placeholder="shadcn" {...field} />
              </ui.FormControl>
              <ui.FormMessage />
            </ui.FormItem>
          )}
        />
        <ui.FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <ui.FormItem>
              <ui.FormLabel>Password</ui.FormLabel>
              <ui.FormControl>
                <Input type="password" placeholder="shadcn" {...field} />
              </ui.FormControl>
              <ui.FormMessage />
            </ui.FormItem>
          )}
        />
        <Button className="px-16" type="submit">
          Submit
        </Button>
      </form>
    </ui.Form>
  );
}
