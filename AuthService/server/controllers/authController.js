import {
  loginUserService,
  registerUserService,
} from "../services/authService.js";

export async function registerUserController(req, res) {
  const { name, surname, email, password } = req.body;
  try {
    const newUser = await registerUserService(name, surname, email, password);
    res
      .status(201)
      .json({ message: "User registered successfully", user: newUser });
  } catch (error) {
    res.status(400).json({ message: error.message });
  }
}

export async function loginUserController(req, res) {
  const { email, password } = req.body;
  try {
    const { token, user } = await loginUserService(email, password);
    res.status(200).json({ message: "Login successful", token, user });
  } catch (error) {
    res.status(400).json({ message: error.message });
  }
}
