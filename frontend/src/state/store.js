import { applyMiddleware, compose, createStore, combineReducers } from 'redux';
import thunk from 'redux-thunk';
import { createLogger } from 'redux-logger';
import { throttle } from 'lodash';

import { watch } from './watch';
import { loadState, saveState } from './localStorage';

// Reducers
import session from 'state/reducers/session';
import profile from 'state/reducers/profile';
import twitterAccounts from 'state/reducers/twitterAccounts';
import campaigns from 'state/reducers/campaigns';
import keywords from 'state/reducers/keywords';
import stats from 'state/reducers/stats';

const logger = createLogger();
const createStoreWithMiddleware = applyMiddleware(thunk, logger)(createStore);
const reducer = combineReducers({
  session,
  profile,
  twitterAccounts,
  campaigns,
  keywords,
  stats
});
const persistedState = loadState();
const store = createStoreWithMiddleware(reducer, persistedState);

// Save oAuth token to localStorage
const saveSessionState = () => {
  console.log("Save session to localStorage");
  saveState({
    session: Object.assign({}, store.getState().session, { isAuthenticating: false })
  });
};
store.subscribe(watch(store.getState, "session")(saveSessionState));

export default store;
