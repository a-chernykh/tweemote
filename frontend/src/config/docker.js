import { getCurrentHost } from 'lib/browser';

const CONFIG = {
  baseURL: getCurrentHost(),
  oauthClientId: 'web',
  oauthClientSecret: 'redacted'
};

export default CONFIG
