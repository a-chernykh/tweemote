import { FETCHED_USER_INFO } from 'constants/actions';

const defaultState = {};

export default function(state = defaultState, action) {
  switch(action.type) {
    case FETCHED_USER_INFO:
      return Object.assign({}, state, {
        email: action.email
      });
  }

  return state;
};
