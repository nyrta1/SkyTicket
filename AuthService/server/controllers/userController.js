import {
  getAllUsersService,
  getUserByIdService,
  addUserService,
  updateUserService,
  deleteUserService,
} from "../services/userService.js";

export async function getAllUsersController() {
  try {
    const users = await getAllUsersService();
    return users;
  } catch (error) {
    console.error(error.message);
    throw error;
  }
}

export async function getUserByIdController(userId) {
  try {
    const user = await getUserByIdService(userId);
    if (!user) {
      console.log(`User with ID ${userId} not found`);
      return null;
    }
    return user;
  } catch (error) {
    console.error("Error fetching user by ID:", error.message);
    throw error;
  }
}

export async function addUserController(name, surname, email, password) {
  try {
    const newUser = await addUserService(name, surname, email, password);
    console.log("User added!");
    return newUser;
  } catch (error) {
    console.error("Error adding user:", error.message);
    throw error;
  }
}

export async function updateUserController(userId, updates) {
  try {
    const updatedUser = await updateUserService(userId, updates);
    console.log("User updated");
    return updatedUser;
  } catch (error) {
    console.error("Error updating user:", error.message);
    throw error;
  }
}

export async function deleteUserController(userId) {
  try {
    const deletedUser = await deleteUserService(userId);
    console.log("User deleted");
    return deletedUser;
  } catch (error) {
    console.error("Error deleting user:", error.message);
    throw error;
  }
}
