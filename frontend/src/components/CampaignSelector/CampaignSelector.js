import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
import { connect } from 'react-redux';
import { Squares } from 'react-activity';

import { fetchCampaigns } from 'state/actions/campaigns';
import { fetchTwitterAccounts } from 'state/actions/twitterAccounts';

class CampaignSelector extends Component {
  constructor(props) {
    super(props);

    this.props.fetchTwitterAccounts();
    this.twitterAccountId = this.props.match.params.id;
    this.props.fetchCampaigns(this.twitterAccountId);
  }

  render() {
    if (this.props.isFetchingCampaigns || this.props.isFetchingAccounts) {
      return <Squares size={40} />;
    } else {
      if (this.props.twitterAccounts.length == 0) {
        return <p>You don't have any Twitter accounts linked.</p>;
      } else {
        let campaignIds = this.props.twitterAccounts[this.twitterAccountId].campaigns;

        if (!campaignIds || campaignIds.length == 0) {
          return <p>You don't have any campaigns created.</p>;
        } else {
          let campaignId = campaignIds[0];

          return <Redirect to={`/dashboard/${this.twitterAccountId}/${campaignId}`} />;
        }
      }
    }
  }
}

const mapStateToProps = (state) => ({
  isFetchingCampaigns: state.campaigns.isFetching,
  isFetchingAccounts: state.twitterAccounts.isFetching,
  twitterAccounts: state.twitterAccounts.items,
  campaigns: state.campaigns.items
});
const mapDispatchToProps = {
  fetchTwitterAccounts,
  fetchCampaigns
};

export default connect(mapStateToProps, mapDispatchToProps)(CampaignSelector);
