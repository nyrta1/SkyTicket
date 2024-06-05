import { createClient } from "@supabase/supabase-js";
import dotenv from "dotenv";
import bcrypt from "bcrypt";
import jwt from "jsonwebtoken";

dotenv.config();

const supabaseUrl = process.env.SUPABASE_URL;
const supabaseKey = process.env.SUPABASE_KEY;
const supabase = createClient(supabaseUrl, supabaseKey);

export async function registerUserService(name, surname, email, password) {
  try {
    const hashedPassword = await bcrypt.hash(password, 10);
    const { data, error } = await supabase
      .from("users")
      .insert([{ name, surname, email, password: hashedPassword }]);
    if (error) {
      throw new Error(error.message);
    }
    return data;
  } catch (error) {
    throw new Error("Error registering user: " + error.message);
  }
}

export async function loginUserService(email, password) {
  try {
    console.log(`Attempting to login with email: ${email}`); // Debug line
    const { data, error } = await supabase
      .from("users")
      .select("*")
      .eq("email", email)
      .single();
    if (error || !data) {
      console.error(
        `User not found or error occurred: ${
          error ? error.message : "No user found"
        }`
      ); // Debug line
      throw new Error("Invalid email or password");
    }

    console.log(`User found: ${JSON.stringify(data)}`); // Debug line
    const isPasswordValid = await bcrypt.compare(password, data.password);
    if (!isPasswordValid) {
      console.error("Password validation failed"); // Debug line
      throw new Error("Invalid email or password");
    }

    const token = jwt.sign({ userId: data.id }, process.env.JWT_SECRET, {
      expiresIn: "1h",
    });
    return { token, user: data };
  } catch (error) {
    console.error(`Error during login: ${error.message}`); // Debug line
    throw new Error("Error logging in: " + error.message);
  }
}
