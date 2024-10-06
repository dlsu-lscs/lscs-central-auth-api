import dotenv from 'dotenv';
import createError from 'http-errors';
import express from 'express';
import path from 'path';
import cookieParser from 'cookie-parser';
import logger from 'morgan';

import session from 'express-session';
import passport from 'passport';
import SQLiteStoreFactory from 'connect-sqlite3';
import pluralize from 'pluralize';
import authRouter from './routes/auth.js';

dotenv.config();

const SQLiteStore = SQLiteStoreFactory(session);
const app = express();

app.locals.pluralize = pluralize;

app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(process.cwd(), 'public')));
app.use(
  session({
    // TODO: get secret from .env
    secret: process.env.SESSION_SECRET || 'keyboard cat',
    resave: false,
    saveUninitialized: false,
    store: new SQLiteStore({ db: 'sessions.db', dir: './const/db' }),
  }),
);

app.use('/', authRouter);

// test: curl http://localhost:3000/
app.listen(process.env.PORT, () => {
  console.log(`API listening on port ${process.env.PORT}`);
});
