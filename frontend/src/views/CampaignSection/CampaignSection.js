import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';
import { Squares } from 'react-activity';

import { fetchCampaigns } from 'state/actions/campaigns';
import { fetchTwitterAccounts } from 'state/actions/twitterAccounts';

import css from './CampaignSection.less';
import PrivateRoute from 'components/PrivateRoute';
import Keywords from 'components/Keywords';
import CampaignDashboard from 'components/CampaignDashboard';
import TwitterAccountLink from 'views/TwitterAccountLink';

class CampaignSection extends Component {
  constructor(props) {
    super(props);

    this.twitterAccountId = this.props.match.params.id;
    this.campaignId = this.props.match.params.campaign_id;

    this.props.fetchTwitterAccounts();
    this.props.fetchCampaigns(this.twitterAccountId);
  }

  render() {
    let account = this.props.twitterAccounts[this.twitterAccountId];
    let campaign = this.props.campaigns[this.campaignId];

    if (account == null || campaign == null) {
      return <Squares size={40} />;
    } else {
      return (
        <div className="row campaign-section">
          <div className="col-xs-3 account-menu">
            <TwitterAccountLink id={account.id} username={account.twitter_username} />
            <ul className="list-unstyled">
              <li><i className="fa fa-bar-chart" aria-hidden="true"></i> <Link to={`${this.props.match.url}`}>Statistics</Link></li>
              <li><i className="fa fa-cogs" aria-hidden="true"></i> <Link to={`${this.props.match.url}/keywords`}>Keywords</Link></li>
            </ul>
          </div>

          <div className="col-xs-9">
            <h2>{campaign.name}</h2>
            <PrivateRoute exact path={this.props.match.path}
                                component={CampaignDashboard}
                                name="Statistics"
                                campaign={campaign} />
            <PrivateRoute exact path={`${this.props.match.path}/keywords`}
                                component={Keywords}
                                name="Keywords"
                                campaign={campaign} />
          </div>
        </div>
      );
    }
  }
}

const mapStateToProps = (state) => ({
  twitterAccounts: state.twitterAccounts.items,
  campaigns: state.campaigns.items
});

const mapDispatchToProps = {
  fetchTwitterAccounts,
  fetchCampaigns
};

export default connect(mapStateToProps, mapDispatchToProps)(CampaignSection);
