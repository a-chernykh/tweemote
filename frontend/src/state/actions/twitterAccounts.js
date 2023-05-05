import { REQUEST_TWITTER_ACCOUNTS,
         FETCHED_TWITTER_ACCOUNTS,
         REQUEST_UNLINK_TWITTER_ACCOUNT,
         UNLINKED_TWITTER_ACCOUNT,
         LINK_TWITTER_ACCOUNT_REDIRECT,
         REQUEST_FINISH_TWITTER_LINK,
         FINISHED_TWITTER_LINK } from 'constants/actions';
import { defaultErrorHandler } from 'lib/ajax';
import { getCurrentHost } from 'lib/browser';

import * as api from 'lib/api';

const requestTwitterAccounts = () => ({
  type: REQUEST_TWITTER_ACCOUNTS
});

const fetchedTwitterAccounts = (accounts) => ({
  type: FETCHED_TWITTER_ACCOUNTS,
  accounts
});

const requestFinishTwitterLink = (appId, oauthToken, oauthVerifier) => ({
  type: REQUEST_FINISH_TWITTER_LINK,
  appId,
  oauthToken,
  oauthVerifier
});

const finishedTwitterLink = (twitterAccountId) => ({
  type: FINISHED_TWITTER_LINK,
  twitterAccountId
});

const requestUnlinkTwitterAccount = (id) => ({
  type: REQUEST_UNLINK_TWITTER_ACCOUNT,
  id
});

const unlinkedTwitterAccount = (id) => ({
  type: UNLINKED_TWITTER_ACCOUNT,
  id
});

export function linkTwitterAccountRedirect() {
  return function(dispatch) {
    dispatch({type: LINK_TWITTER_ACCOUNT_REDIRECT});
    let cb = getCurrentHost() + "/link/twitter/callback";
    api.linkTwitterAccount(cb).catch(defaultErrorHandler);
  };
};

export function finishTwitterLink(appId, oauthToken, oauthVerifier) {
  return function(dispatch) {
    dispatch(requestFinishTwitterLink(appId, oauthToken, oauthVerifier));

    api.finishTwitterLink(appId, oauthToken, oauthVerifier)
      .then((a) => {
        dispatch(finishedTwitterLink(a.id));
      }, defaultErrorHandler);
  };
};

export function fetchTwitterAccounts() {
  return function(dispatch) {
    dispatch(requestTwitterAccounts());

    api.fetchTwitterAccounts()
      .then(accounts => {
        dispatch(fetchedTwitterAccounts(accounts));
      }, defaultErrorHandler);
  };
};

export function unlinkTwitterAccount(id) {
  return function(dispatch) {
    dispatch(requestUnlinkTwitterAccount(id));

    api.unlinkTwitterAccount(id)
      .then(() => {
        dispatch(unlinkedTwitterAccount(id));
      }, defaultErrorHandler);
  };
};
