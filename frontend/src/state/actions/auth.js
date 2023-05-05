import { postAuth, post, defaultErrorHandler } from 'lib/ajax';
import { SIGNIN, SIGNIN_SUCCESSFUL, SIGNIN_FAILURE,
         REFRESH, REFRESH_SUCCESSFUL, REFRESH_FAILURE,
         SIGNUP, SIGNUP_SUCCESSFUL, SIGNUP_FAILURE,
         SIGNOUT } from 'constants/actions';
import { notifier } from 'lib/notifier';
import * as api from 'lib/api';

export const signinRequest = (email, password) => ({
  type: SIGNIN
});

export const signinSuccessful = (accessToken, refreshToken, expiresIn) => ({
  type: SIGNIN_SUCCESSFUL,
  accessToken,
  refreshToken,
  expiresIn
});

export const signinFailure = (error) => ({
  type: SIGNIN_FAILURE,
  error
});

export const refreshRequest = () => ({
  type: REFRESH
});

export const refreshSuccessful = (accessToken, refreshToken, expiresIn) => ({
  type: REFRESH_SUCCESSFUL,
  accessToken,
  refreshToken,
  expiresIn
});

export const refreshFailure = (error) => ({
  type: REFRESH_FAILURE,
  error
});

export const signupRequest = (email, password) => ({
  type: SIGNUP
});

export const signupSuccessful = () => ({
  type: SIGNUP_SUCCESSFUL
});

export const signupFailure = (error) => ({
  type: SIGNUP_FAILURE,
  error
});

export const signout = () => ({
  type: SIGNOUT
});

export function signinUser(email, password) {
  return function(dispatch) {
    dispatch(signinRequest());

    let params = {
      grant_type: "password",
      username: email,
      password: password,
      scope: "write"
    };

    postAuth(params)
    .then(r => {
      notifier.info("Signed in successfully");
      dispatch(signinSuccessful(r.data.access_token,
                                r.data.refresh_token,
                                r.data.expires_in));
    }, e => {
      let error;

      if (e.response && e.response.data && e.response.data.error) {
        if (e.response.data.error == 'access_denied') {
          error = "Please verify your e-mail and password.";
        } else {
          error = e.response.data.error;
        }
      } else {
        error = 'Unexpected server error';
        console.log(e);
      }

      notifier.error(error, "Unable to sign in", {autoDismiss: 5});
      dispatch(signinFailure(error));
    });
  };
};

export function signupUser(email, password, passwordConfirmation) {
  return function(dispatch) {
    dispatch(signupRequest());

    post('/api/v1/users', {
      email: email,
      password: password,
      password_confirmation: passwordConfirmation
    })
    .then(r => {
      notifier.info("Signed up successfully");
      dispatch(signupSuccessful());
    }, e => {
      defaultErrorHandler(e);
      let err = "Network error";
      if (e.response && e.response.data) {
        err = e.response.data.error;
      }
      dispatch(signupFailure(err));
    });
  };
};

export function signoutUser() {
  return function(dispatch) {
    // TODO send token invalidation request here
    dispatch(signout());
  };
};
