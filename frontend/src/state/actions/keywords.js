import { REQUEST_KEYWORDS,
         FETCHED_KEYWORDS,
         REQUEST_ADD_KEYWORD,
         ADD_KEYWORD_SUCCESSFUL,
         ADD_KEYWORD_FAILURE,
         REQUEST_DELETE_KEYWORD,
         DELETE_KEYWORD_SUCCESSFUL,
         DELETE_KEYWORD_FAILURE } from 'constants/actions';
import { defaultErrorHandler } from 'lib/ajax';
import * as api from 'lib/api';

const requestKeywords = (campaignId) => ({
  type: REQUEST_KEYWORDS,
  campaignId
});

const fetchedKeywords = (campaignId, keywords) => ({
  type: FETCHED_KEYWORDS,
  campaignId,
  keywords
});

const requestAddKeyword = (campaignId, keyword) => ({
  type: REQUEST_ADD_KEYWORD,
  campaignId,
  keyword
});

const addKeywordSuccessful = (campaignId, keyword) => ({
  type: ADD_KEYWORD_SUCCESSFUL,
  campaignId,
  keyword
});

const addKeywordFailed = (campaignId, keyword) => ({
  type: ADD_KEYWORD_FAILURE,
  campaignId,
  keyword
});

const requestDeleteKeyword = (campaignId, keywordId) => ({
  type: REQUEST_DELETE_KEYWORD,
  campaignId,
  keywordId
});

const deleteKeywordSuccessful = (campaignId, keywordId) => ({
  type: DELETE_KEYWORD_SUCCESSFUL,
  campaignId,
  keywordId
});

const deleteKeywordFailed = (campaignId, keywordId) => ({
  type: DELETE_KEYWORD_FAILURE,
  campaignId,
  keywordId
});

export function fetchKeywords(campaignId) {
  return function(dispatch) {
    dispatch(requestKeywords(campaignId));

    api.fetchKeywords(campaignId)
      .then(keywords => {
        dispatch(fetchedKeywords(campaignId, keywords));
      }, defaultErrorHandler);
  };
};

export function addKeyword(campaignId, keyword) {
  return function(dispatch) {
    dispatch(requestAddKeyword(campaignId, keyword));

    api.addKeyword(campaignId, keyword)
      .then(keyword => {
        dispatch(addKeywordSuccessful(campaignId, keyword));
      }, r => {
        dispatch(addKeywordFailed(campaignId, keyword));
        defaultErrorHandler(r);
      });
  };
};

export function deleteKeyword(campaignId, keywordId) {
  return function(dispatch) {
    dispatch(requestDeleteKeyword(campaignId, keywordId));

    api.deleteKeyword(campaignId, keywordId)
      .then(() => {
        dispatch(deleteKeywordSuccessful(campaignId, keywordId));
      }, r => {
        dispatch(deleteKeywordFailed(campaignId, keywordId));
        defaultErrorHandler(r);
      });
  };
};
