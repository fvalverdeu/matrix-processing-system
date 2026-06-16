require('dotenv').config();

const express = require('express');
const config = require('../config');

const app = express();

app.use(express.json());

app.get('/health', (req, res) => {
  res.json({ status: 'ok', service: 'statistics-service-node' });
});

app.listen(config.port, () => {
  console.log(`Statistics service listening on port ${config.port}`);
});
