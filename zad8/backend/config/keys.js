require('dotenv').config();

module.exports = {
    googleClientID: process.env.GOOGLE_CLIENT_ID || 'YOUR_GOOGLE_CLIENT_ID',
    googleClientSecret: process.env.GOOGLE_CLIENT_SECRET || 'YOUR_GOOGLE_CLIENT_SECRET',
    jwtSecret: process.env.JWT_SECRET || 'YOUR_JWT_SECRET',
    mongoURI: process.env.MONGO_URI || 'YOUR_MONGO_URI'
  };