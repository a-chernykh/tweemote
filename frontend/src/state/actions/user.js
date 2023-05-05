import { FETCHED_USER_INFO } from 'constants/actions';
import * as api from 'lib/api';

export const fetchedUserInfo = (email) => ({
  type: FETCHED_USER_INFO,
  email
});

export function fetchUserInfo() {
  return function(dispatch) {
    api.fetchUserInfo().then(userInfo => {
      dispatch(fetchedUserInfo(userInfo.email));
    });
  };
};
