import axios from 'axios';

import config from 'config';
import store from 'state/store';
import { refreshRequest, refreshSuccessful, refreshFailure } from 'state/actions/auth';
import { notifier } from './notifier';

const axiosInstance = axios.create({
  baseURL: config.baseURL,
  timeout: 2000
});

// Append `Authorization: Bearer <Access token>` header
axiosInstance.interceptors.request.use(config => {
  if (store.getState().session.isAuthenticated) {
    config.headers['Authorization'] = `Bearer ${store.getState().session.accessToken}`;
  }
  return config;
}, error => {
  return Promise.reject(error);
});

let refreshPromise = null;
const refreshToken = () => {
  if (!refreshPromise) {
    refreshPromise = new Promise((resolve, reject) => {
      store.dispatch(refreshRequest());

      let refreshToken = store.getState().session.refreshToken;

      if (refreshToken) {
        postAuth({
          grant_type: "refresh_token",
          refresh_token: refreshToken
        }).then(r => {
          store.dispatch(refreshSuccessful(
            r.data.access_token,
            r.data.refresh_token,
            r.data.expires_in
          ));
          resolve(r);
        }).catch(e => {
          console.log(`Failed to refresh token because of error: ${e}`);
          store.dispatch(refreshFailure());
          reject(e);
        });
      } else {
        console.log("Failed to refresh token because refresh token is not set");
        store.dispatch(refreshFailure());
        reject(e);
      }
    });

    refreshPromise.then(r => {
      refreshPromise = null;
    });
  }

  return refreshPromise;
};

// Refresh access token
axiosInstance.interceptors.response.use(r => {
  return r;
}, e => {
  let r = e.response;

  // Sometimes response is not set (example: when request is cancelled)
  if (r != null) {
    let originalConfig = r.config;

    if (!originalConfig.retried && r.status == 401) {
      if (r.data.error == "invalid_grant" || r.data.error == "invalid_request") {
        // "invalid_request" response will be received when refresh request succeded before this request was started, therefore token used in this request was already refreshed

        return new Promise((resolve, reject) => {
          refreshToken().then(r => {
            console.log(`Re-trying request with ${r.data.access_token}`);

            originalConfig.headers['Authorization'] = `Bearer ${r.data.access_token}`;
            originalConfig.retried = true
            axiosInstance.request(originalConfig).then(r => {
              resolve(r);
            }).catch(e => {
              reject(e);
            });
          }).catch(e => {
            reject(e);
          });
        });
      } else {
        // TODO more generic error text
        notifier.error(r.data.error_description);
        console.log("Failed to refresh token because of different HTTP error");
        store.dispatch(refreshFailure());
      }
    }
  }

  return Promise.reject(e);
});

function postAuth(params) {
  let urlParams = new URLSearchParams();
  for (let k in params) {
    urlParams.append(k, params[k]);
  }

  let reqConfig = {
    auth: {
      username: config.oauthClientId,
      password: config.oauthClientSecret
    }
  };

  return post('/api/v1/auth', urlParams, reqConfig);
}

function post(url, data, config={}) {
  return axiosInstance.post(url, data, config);
}

function get(url, config={}) {
  return axiosInstance.get(url, config);
}

function del(url, config={}) {
  return axiosInstance.delete(url, config);
}

function defaultErrorHandler(e) {
  let msg = `Unexpected error: ${e.message}`;

  if (e.response && e.response.data && e.response.data.error_description) {
    msg = e.response.data.error_description;
  }

  notifier.error(msg);
}

export { post, postAuth, get, del, defaultErrorHandler };
