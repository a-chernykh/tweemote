import { REQUEST_CAMPAIGNS,
         FETCHED_CAMPAIGNS } from 'constants/actions';
import { defaultErrorHandler } from 'lib/ajax';
import * as api from 'lib/api';

const requestCampaigns = (twitterAccountId) => ({
  type: REQUEST_CAMPAIGNS,
  twitterAccountId
});

const fetchedCampaigns = (twitterAccountId, campaigns) => ({
  type: FETCHED_CAMPAIGNS,
  twitterAccountId,
  campaigns
});

export function fetchCampaigns(twitterAccountId) {
  return function(dispatch) {
    dispatch(requestCampaigns(twitterAccountId));

    api.fetchCampaigns(twitterAccountId)
      .then(campaigns => {
        dispatch(fetchedCampaigns(twitterAccountId, campaigns));
      }, defaultErrorHandler);
  };
};
