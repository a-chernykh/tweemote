import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
import { connect } from 'react-redux';
import { Squares } from 'react-activity';

import { fetchTwitterAccounts } from 'state/actions/twitterAccounts';
import CleanDashboard from 'views/CleanDashboard';
import Dashboard from 'views/Dashboard';
import AccountSelector from 'views/AccountSelector';

class UserDashboard extends Component {
  componentDidMount() {
    this.props.fetchTwitterAccounts();
  }

  render() {
    if (this.props.isFetching) {
      return <Squares size={40} />;
    } else {
      let accountIds = Object.keys(this.props.accounts);
      if (accountIds.length > 0) {
        if (accountIds.length == 1) {
          let acctId = accountIds[0];
          return <Redirect to={`/dashboard/${acctId}`} />;
        } else {
          return <AccountSelector accounts={Object.values(this.props.accounts)} />
        }
      } else {
        return <CleanDashboard />;
      }
    }
  }
}

const mapStateToProps = (state) => ({
  isFetching: state.twitterAccounts.isFetching,
  accounts: state.twitterAccounts.items
});
const mapDispatchToProps = {
  fetchTwitterAccounts
};

export default connect(mapStateToProps, mapDispatchToProps)(UserDashboard);
