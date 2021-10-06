export default {
  port: process.env.RIMGU_PORT || 8080,
  host: process.env.RIMGU_HOST || 'localhost',
  address: process.env.RIMGU_ADDRESS || '127.0.0.1',
  http_proxy: process.env.RIMGU_HTTP_PROXY || null,
  https_proxy: process.env.RIMGU_HTTPS_PROXY || null,
  imgur_client_id: process.env.RIMGU_IMGUR_CLIENT_ID || null,
  disable_comments: process.env.RIMGU_DISABLE_COMMENTS === 'true',
};
