const app = require('./app');
const mongoose = require('mongoose');
const keys = require('./config/keys');

mongoose.connect('mongodb://localhost:27017/oauth_demo', { useNewUrlParser: true, useUnifiedTopology: true })
  .then(() => app.listen(5000, () => console.log('Server running on port 5000')))
  .catch(err => console.error(err));
