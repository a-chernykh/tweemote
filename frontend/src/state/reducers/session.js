import { SIGNIN, SIGNIN_SUCCESSFUL, SIGNIN_FAILURE,
         REFRESH, REFRESH_SUCCESSFUL, REFRESH_FAILURE,
         SIGNUP, SIGNUP_SUCCESSFUL, SIGNUP_FAILURE,
         SIGNOUT } from 'constants/actions';

const defaultState = {
  isAuthenticating: false,
  isAuthenticated: false,
  isRegistering: false,
  isRegistered: false
};

export default function(state = defaultState, action) {
  switch(action.type) {
    case SIGNIN:
      return Object.assign({}, state, { isAuthenticating: true, isAuthenticated: false });
    case REFRESH:
      return state;
    case SIGNIN_SUCCESSFUL:
    case REFRESH_SUCCESSFUL:
      return Object.assign({}, state, {
        isAuthenticating: false,
        isAuthenticated: true,
        accessToken: action.accessToken,
        refreshToken: action.refreshToken,
        expiresIn: action.expiresIn
      });
    case SIGNIN_FAILURE:
    case REFRESH_FAILURE:
      return Object.assign({}, state, {
        isAuthenticating: false,
        isAuthenticated: false,
        isRegistered: false,
        accessToken: undefined,
        refreshToken: undefined,
        expiresIn: undefined
      });

    case SIGNUP:
      return Object.assign({}, state, { isRegistering: true });
    case SIGNUP_SUCCESSFUL:
      return Object.assign({}, state, { isRegistering: false, isRegistered: true });
    case SIGNUP_FAILURE:
      return Object.assign({}, state, { isRegistering: false, isRegistered: false });

    case SIGNOUT:
      return Object.assign({}, state, defaultState);
  }

  return state;
};
