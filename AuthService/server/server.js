import pgdb from "pg";
import dotenv from "dotenv";

const { Client } = pgdb;
dotenv.config();

const client = new Client({
  host: process.env.POSTGRES_URL,
  port: process.env.POSTGRES_PORT,
  user: process.env.POSTGRES_USERNAME,
  password: process.env.POSTGRES_PASSWORD,
  database: process.env.POSTGRES_DATABASE,
});

client
  .connect()
  .then(() => console.log("Connected to PostgreSQL"))
  .catch((err) => console.error("Connection error", err.stack));
