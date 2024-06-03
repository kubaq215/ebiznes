const express = require('express');
const mongoose = require('mongoose');
const User = mongoose.model('User');
const jwt = require('jsonwebtoken');
const keys = require('../config/keys');

const router = express.Router();

router.post('/', async (req, res) => {
  const { email, password } = req.body;

  if (!email || !password) {
    return res.status(400).send('All fields are required');
  }

  const user = await User.findOne({ email, password });
  if (!user) {
    return res.status(400).send('Invalid credentials');
  }

  const token = jwt.sign({ userId: user.id }, keys.jwtSecret, { expiresIn: '1h' });
  res.send({ token });
});

module.exports = router;
