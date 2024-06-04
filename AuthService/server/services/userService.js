import { createClient } from "@supabase/supabase-js";
import dotenv from "dotenv";

dotenv.config();

const supabaseUrl = process.env.SUPABASE_URL;
const supabaseKey = process.env.SUPABASE_KEY;
const supabase = createClient(supabaseUrl, supabaseKey);

export async function getAllUsersService() {
  try {
    const { data, error } = await supabase.from("users").select("*");
    if (error) {
      throw new Error(error.message);
    }
    return data;
  } catch (error) {
    throw new Error("Error fetching users: " + error.message);
  }
}

export async function getUserByIdService(userId) {
  try {
    const { data, error } = await supabase
      .from("users")
      .select("*")
      .eq("id", userId)
      .single();
    if (error) {
      throw new Error(error.message);
    }
    return data;
  } catch (error) {
    throw new Error("Error fetching user by ID: " + error.message);
  }
}

export async function addUserService(name, surname, email, password) {
  try {
    const { data, error } = await supabase
      .from("users")
      .insert([{ name, surname, email, password }]);
    if (error) {
      throw new Error(error.message);
    }
    return data;
  } catch (error) {
    throw new Error("Error adding user: " + error.message);
  }
}

export async function updateUserService(userId, updates) {
  try {
    const { data, error } = await supabase
      .from("users")
      .update(updates)
      .eq("id", userId);
    if (error) {
      throw new Error(error.message);
    }
    return data;
  } catch (error) {
    throw new Error("Error updating user: " + error.message);
  }
}

export async function deleteUserService(userId) {
  try {
    const { data, error } = await supabase
      .from("users")
      .delete()
      .eq("id", userId);
    if (error) {
      throw new Error(error.message);
    }
    return data;
  } catch (error) {
    throw new Error("Error deleting user: " + error.message);
  }
}
