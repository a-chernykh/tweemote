import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
import { connect } from 'react-redux';
import { Squares } from 'react-activity';

import { finishTwitterLink } from 'state/actions/twitterAccounts';
import { getParameterByName } from 'lib/browser';

class LinkTwitterCallback extends Component {
  componentWillMount() {
    let appId = getParameterByName("twitter_application_id"),
        oauthToken = getParameterByName("oauth_token"),
        oauthVerifier = getParameterByName("oauth_verifier");
    this.props.finishTwitterLink(appId, oauthToken, oauthVerifier);
  }

  render() {
    if (this.props.isFinishingLink == undefined || this.props.isFinishingLink) {
      return (
        <Squares size={40} />
      );
    } else {
      return <Redirect to="/dashboard" />;
    }
  }
}

const mapStateToProps = (state) => ({
  isFinishingLink: state.twitterAccounts.isFinishingLink
});
const mapDispatchToProps = {
  finishTwitterLink
};

export default connect(mapStateToProps, mapDispatchToProps)(LinkTwitterCallback);
