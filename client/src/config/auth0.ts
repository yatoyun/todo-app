export const auth0Config = {
  domain: process.env.REACT_APP_AUTH0_DOMAIN || 'YOUR_AUTH0_DOMAIN',
  clientId: process.env.REACT_APP_AUTH0_CLIENT_ID || 'YOUR_AUTH0_CLIENT_ID',
  redirectUri: window.location.origin,
  audience: process.env.REACT_APP_AUTH0_AUDIENCE || 'YOUR_AUTH0_AUDIENCE',
};
