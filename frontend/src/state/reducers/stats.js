import { REQUEST_STATS,
         FETCHED_STATS } from 'constants/actions';

const defaultState = {
  items:      {},
  isFetching: false,
  isFetched:  false
};

export default function(state = defaultState, action) {
  switch(action.type) {

    case REQUEST_STATS: {
      return Object.assign({}, state, { isFetching: true, isFetched: false });
    }

    case FETCHED_STATS: {
      let newState = { ...state, isFetching: false, isFetched: true };
      newState.items = { ...newState.items, [action.campaignId]: action.stats };
      return newState;
    }

  }

  return state;
};
