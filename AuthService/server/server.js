import express from "express";
import dotenv from "dotenv";
import bodyParser from "body-parser";
import {
  registerUserController,
  loginUserController,
} from "./controllers/authController.js";

dotenv.config();

const app = express();
app.use(bodyParser.json());

app.post("/register", registerUserController);
app.post("/login", loginUserController);

const port = process.env.PORT || 1111;
app.listen(port, () => {
  console.log(`Server running on port ${port}`);
});
