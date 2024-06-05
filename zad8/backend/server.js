const mongoose = require('mongoose');
require('./models/User');
const app = require('./app');
const keys = require('./config/keys');

mongoose.connect(keys.mongoURI)
  .then(() => app.listen(5000, () => console.log('Server running on port 5000')))
  .catch(err => console.error(err));
