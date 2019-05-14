/* eslint-disable import/no-extraneous-dependencies */
const http = require('http');
const httpProxy = require('http-proxy');

const proxy = httpProxy.createProxyServer({ secure: false });

const isProxied = req => req.url.includes('/round');

proxy.on('proxyReq', (proxyReq, req, res) => {
  if (isProxied(req)) {
    proxyReq.setHeader('Host', 'localhost:8000');
  }

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
  if (req.method === 'OPTIONS') {
    res.send(200);
    res.end();
  }
});

function proxyRequest(req, res, target) {
  proxy.web(req, res, { target }, () => {
    const message = `Failed to connect to ${target}`;
    res.statusCode = 500;
    res.end(JSON.stringify({ message }));
  });
}

const server = http.createServer((req, res) => {
  if (isProxied(req)) {
    proxyRequest(req, res, 'http://localhost:8080');
  } else {
    proxyRequest(req, res, 'http://127.0.0.1:3000');
  }
});

server.listen(9000);
