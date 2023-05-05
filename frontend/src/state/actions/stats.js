import { REQUEST_STATS,
         FETCHED_STATS } from 'constants/actions';
import { defaultErrorHandler } from 'lib/ajax';
import * as api from 'lib/api';

const requestStats = (campaignId) => ({
  type: REQUEST_STATS,
  campaignId
});

const fetchedStats = (campaignId, stats) => ({
  type: FETCHED_STATS,
  campaignId,
  stats
});

export function fetchStats(campaignId) {
  return function(dispatch) {
    dispatch(requestStats(campaignId));

    api.fetchStats(campaignId)
      .then(stats => {
        dispatch(fetchedStats(campaignId, stats));
      }, defaultErrorHandler);
  };
};
