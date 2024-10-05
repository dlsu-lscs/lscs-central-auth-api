require("dotenv").config();

const createError = require("http-errors");
const express = require("express");
const path = require("path");
const cookieParser = require("cookie-parser");
const logger = require("morgan");

const session = require("express-session");
const passport = require("passport");

const SQLiteStore = require("connect-sqlite3")(session);

const authRouter = require("./routes/auth");

const app = express();

app.locals.pluralize = require("pluralize");

app.use(logger("dev"));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, "public")));
app.use(
    session({
        // TODO: get secret from .env
        secret: "keyboard cat",
        resave: false,
        saveUninitialized: false,
        store: new SQLiteStore({ db: "sessions.db", dir: "./const/db" }),
    }),
);

app.use("/", authRouter);

module.exports = app;
