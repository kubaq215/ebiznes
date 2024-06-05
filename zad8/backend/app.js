const express = require('express');
const mongoose = require('mongoose');
const session = require('express-session');
const passport = require('passport');
const cors = require('cors');
const authRoutes = require('./routes/authRoutes');
const registerRoutes = require('./routes/registerRoutes');
const loginRoutes = require('./routes/loginRoutes');
require('./models/User');
require('./services/passport');

const app = express();
app.use(express.json());
app.use(session({ secret: 'SECRET', resave: false, saveUninitialized: true }));
app.use(passport.initialize());
app.use(passport.session());

// Add CORS middleware
const corsOptions = {
  origin: 'http://localhost:3000', // frontend URL
  credentials: true,
};
app.use(cors(corsOptions));

app.use('/auth', authRoutes);
app.use('/register', registerRoutes);
app.use('/login', loginRoutes);

module.exports = app;
