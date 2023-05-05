import React, { Component } from 'react';
import { Squares } from 'react-activity';
import { connect } from 'react-redux';

import { linkTwitterAccountRedirect } from 'state/actions/twitterAccounts';

class LinkTwitter extends Component {
  componentDidMount() {
    this.props.linkTwitterAccountRedirect();
  }

  render() {
    return (
      <div>
        <p>Redirecting to Twitter website, please wait...</p>
        <Squares size={40} />
      </div>
    );
  }
}

const mapStateToProps = (state) => ({});
const mapDispatchToProps = {
  linkTwitterAccountRedirect
};

export default connect(mapStateToProps, mapDispatchToProps)(LinkTwitter);
