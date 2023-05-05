import { REQUEST_CAMPAIGNS,
         FETCHED_CAMPAIGNS } from 'constants/actions';

const defaultState = {
  items:      {},
  isFetching: false
};

export default function(state = defaultState, action) {
  switch(action.type) {
    case REQUEST_CAMPAIGNS:
      return Object.assign({}, state, { isFetching: true });

    case FETCHED_CAMPAIGNS:
      let newState = { ...state, isFetching: false };
      for (let campaign of action.campaigns) {
        newState.items = {
          ...newState.items,
          [campaign.id]: campaign
        };
      };
      return newState;
  }

  return state;
};
