import { REQUEST_TWITTER_ACCOUNTS,
         FETCHED_TWITTER_ACCOUNTS,
         LINK_TWITTER_ACCOUNT_REDIRECT,
         REQUEST_FINISH_TWITTER_LINK,
         FINISHED_TWITTER_LINK,
         UNLINKED_TWITTER_ACCOUNT,
         FETCHED_CAMPAIGNS } from 'constants/actions';

const defaultState = {
  items: [],
  isFetching: false,
  isFinishingLink: false
};

export default function(state = defaultState, action) {
  switch(action.type) {
    case REQUEST_TWITTER_ACCOUNTS: {
      return {...state, isFetching: true };
    }

    case FETCHED_TWITTER_ACCOUNTS: {
      let newState = { ...state, isFetching: false };
      for (let account of action.accounts) {
        newState.items = {
          ...newState.items,
          [account.id]: account
        };
      }
      return newState;
    }

    case LINK_TWITTER_ACCOUNT_REDIRECT:
    case REQUEST_FINISH_TWITTER_LINK: {
      return {...state, isFinishingLink: true };
    }

    case FINISHED_TWITTER_LINK: {
      return {...state, isFinishingLink: false };
    }

    case UNLINKED_TWITTER_ACCOUNT: {
      let newState = { ...state };
      delete newState.items[action.id];
      return newState;
    }

    case FETCHED_CAMPAIGNS: {
      return {...state,
              items: {
                ...state.items,
                [action.twitterAccountId]: {
                  ...state.items[action.twitterAccountId],
                  campaigns: action.campaigns.map(c => c.id)
                }
              }
            };
    }
  }

  return state;
};
