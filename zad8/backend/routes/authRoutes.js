const express = require('express');
const passport = require('passport');
const jwt = require('jsonwebtoken');
const keys = require('../config/keys');

const router = express.Router();

router.get('/google', passport.authenticate('google', { scope: ['profile', 'email'] }));

router.get('/google/callback', passport.authenticate('google'), (req, res) => {
  console.log(req.user);
  const token = jwt.sign({ userId: req.user.id }, keys.jwtSecret, { expiresIn: '300s' });
  res.redirect(`http://172-104-249-16.ip.linodeusercontent.com?token=${token}`);
});

module.exports = router;
