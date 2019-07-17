/* eslint-disable import/no-extraneous-dependencies */
const express = require('express');

const app = express();

function applyCORS(req, res) {
  if (req.headers.origin) {
    res.setHeader('access-control-allow-origin', req.headers.origin);
    res.setHeader('access-control-allow-credentials', 'true');
    res.setHeader('access-control-max-age', 60 * 60 * 24 * 30);
  }
  if (req.headers['access-control-request-method']) {
    res.setHeader('access-control-allow-methods', req.headers['access-control-request-method']);
  }
  if (req.headers['access-control-request-headers']) {
    res.setHeader('access-control-allow-headers', req.headers['access-control-request-headers']);
  }
}

app.options('/*', (req, res) => {
  applyCORS(req, res);
  res.sendStatus(200);
});
app.post('/round', (req, res) => {
  applyCORS(req, res);
  res.send({
    token: 'tom@beer.com',
    roundUrl: 'aUrl',
    participants: ['Bob'],
  });
});
app.post('/round/:roundId', (req, res) => {
  applyCORS(req, res);
  res.send({
    token: 'tom@beer.com',
    roundUrl: 'aUrl',
    participants: ['Bob'],
  });
});
app.get('/round', (req, res) => {
  applyCORS(req, res);
  res.send({
    url: 'aUrl',
    participants: [{
      uuid: 'ce118b6e-d8e1-11e7-9296-cec278b6b50a',
      name: 'Tom',
      roundCount: 10,
    }],
    currentCandidate: {
      uuid: 'ce118b6e-d8e1-11e7-9296-cec278b6b50a',
      name: 'Tom',
      roundCount: 10,
    },
  });
});
app.put('/round', (req, res) => {
  applyCORS(req, res);
  res.send({
    url: 'aUrl',
    participants: [
      {
        uuid: '5559be5c-2d73-446b-a3f8-da14d7c7f5a6',
        name: 'Geoff',
        roundCount: 11,
      }],
    currentCandidate: {
      uuid: '5559be5c-2d73-446b-a3f8-da14d7c7f5a6',
      name: 'Geoff',
      roundCount: 11,
    },
  });
});

app.listen(8080);
