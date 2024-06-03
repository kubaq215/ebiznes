const express = require('express');
const passport = require('passport');
const jwt = require('jsonwebtoken');
const keys = require('../config/keys');

const router = express.Router();

router.get('/google', passport.authenticate('google', { scope: ['profile', 'email'] }));

router.get('/google/callback', passport.authenticate('google'), (req, res) => {
  const token = jwt.sign({ userId: req.user.id }, keys.jwtSecret, { expiresIn: '1h' });
  res.redirect(`http://localhost:3000?token=${token}`);
});

module.exports = router;
