import { get, post, del } from 'lib/ajax';
import config from 'config';

//
// User
//
export const fetchUserInfo = () => {
  return new Promise((resolve, reject) => {
    get("/api/v1/users/me").then(r => {
      resolve({
        email: r.data.email
      });
    }).catch(e => {
      reject(e);
    });
  });
};

//
// Twitter accounts
//
export const linkTwitterAccount = (cb) => {
  let url = config.baseURL + "/api/v1/twitter_accounts/link?callback=" + cb;

  return new Promise((resolve, reject) => {
    get(url).then(r => {
      window.location = r.data.redirect_url;
    }).catch(e => {
      reject(e);
    });
  });
};

export const finishTwitterLink = (appId, oauthToken, oauthVerifier) => {
  return new Promise((resolve, reject) => {
    get("/api/v1/twitter_accounts/callback", { params: {
      twitter_application_id: appId,
      oauth_token: oauthToken,
      oauth_verifier: oauthVerifier
    }}).then(r => {
      resolve(r.data);
    }).catch(e => {
      reject(e);
    });
  });
};

export const fetchTwitterAccounts = () => {
  return new Promise((resolve, reject) => {
    get("/api/v1/twitter_accounts").then(r => {
      resolve(r.data.twitter_accounts);
    }).catch(e => {
      reject(e);
    });
  });
};

export const unlinkTwitterAccount = (id) => {
  return new Promise((resolve, reject) => {
    del(`/api/v1/twitter_accounts/${id}`).then(r => {
      resolve();
    }).catch(e => {
      reject(e);
    });
  });
};

//
// Campaigns
//
export const fetchCampaigns = (twitterAccountId) => {
  return new Promise((resolve, reject) => {
    get(`/api/v1/twitter_accounts/${twitterAccountId}/campaigns`).then(r => {
      resolve(r.data.campaigns);
    }).catch(e => {
      reject(e);
    });
  });
};

//
// Keywords
//
export const fetchKeywords = (campaignId) => {
  return new Promise((resolve, reject) => {
    get(`/api/v1/campaigns/${campaignId}/keywords`).then(r => {
      resolve(r.data.keywords);
    }).catch(e => {
      reject(e);
    });
  });
};

export const addKeyword = (campaignId, keyword) => {
  return new Promise((resolve, reject) => {
    post(`/api/v1/campaigns/${campaignId}/keywords`, { keyword: keyword }).then(r => {
      resolve(r.data.keyword);
    }).catch(e => {
      reject(e);
    });
  });
};

export const deleteKeyword = (campaignId, keywordId) => {
  return new Promise((resolve, reject) => {
    del(`/api/v1/campaigns/${campaignId}/keywords/${keywordId}`).then(() => {
      resolve();
    }).catch(e => {
      reject(e);
    });
  });
};

//
// Stats
//
export const fetchStats = (campaignId) => {
  return new Promise((resolve, reject) => {
    get(`/api/v1/campaigns/${campaignId}/stats`).then(r => {
      resolve(r.data.stats);
    }).catch(e => {
      reject(e);
    });
  });
};
