import { REQUEST_KEYWORDS,
         FETCHED_KEYWORDS,
         REQUEST_ADD_KEYWORD,
         ADD_KEYWORD_SUCCESSFUL,
         ADD_KEYWORD_FAILURE,
         REQUEST_DELETE_KEYWORD,
         DELETE_KEYWORD_SUCCESSFUL,
         DELETE_KEYWORD_FAILURE } from 'constants/actions';

const defaultState = {
  items:      {},
  isFetching: false,
  isFetched:  false,
  isAdding:   false
};

export default function(state = defaultState, action) {
  switch(action.type) {
    case REQUEST_KEYWORDS: {
      return Object.assign({}, state, { isFetching: true, isFetched: false });
    }

    case FETCHED_KEYWORDS: {
      let newState = { ...state, isFetching: false, isFetched: true };
      for (let keyword of action.keywords) {
        newState.items = {
          ...newState.items,
          [keyword.id]: keyword
        };
      };
      return newState;
    }

    case REQUEST_ADD_KEYWORD: {
      return Object.assign({}, state, { isAdding: true });
    }

    case ADD_KEYWORD_SUCCESSFUL: {
      let newState = { ...state, isAdding: false };
      newState.items = { ...newState.items, [action.keyword.id]: { ...action.keyword, isNew: true } };
      return newState;
    }

    case ADD_KEYWORD_FAILURE: {
      return Object.assign({}, state, { isAdding: false });
    }

    case REQUEST_DELETE_KEYWORD: {
      let newState = { ...state, items: {
          ...state.items,
          [action.keywordId]: {
            ...state.items[action.keywordId],
            isDeleting: true
          }
        }
      };
      return newState;
    }

    case DELETE_KEYWORD_SUCCESSFUL: {
      let newState = {
        ...state,
        items: { ...state.items }
      };
      delete newState.items[action.keywordId];
      return newState;
    }

    case DELETE_KEYWORD_FAILURE: {
      let newState = { ...state, items: {
          ...state.items,
          [action.keywordId]: {
            ...state.items[action.keywordId],
            isDeleting: false
          }
        }
      };
      return newState;
    }

  }

  return state;
};
