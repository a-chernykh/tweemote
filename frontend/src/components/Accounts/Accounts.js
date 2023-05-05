import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Squares } from 'react-activity';

import TwitterAccountLink from 'views/TwitterAccountLink';
import DropdownButton from 'views/DropdownButton';
import DropdownAction from 'views/DropdownAction';
import CleanDashboard from 'views/CleanDashboard';
import { fetchTwitterAccounts, unlinkTwitterAccount } from 'state/actions/twitterAccounts';
import css from './Accounts.less';

class Accounts extends Component {
  componentDidMount() {
    if (Object.keys(this.props.accounts).length == 0) {
      this.props.fetchTwitterAccounts();
    }
  }

  render() {
    if (this.props.isFetching) {
      return <Squares size={40} />;
    } else {
      let accounts = [];
      for(let id in this.props.accounts) {
        let a = this.props.accounts[id];

        accounts.push(
          <li key={a.twitter_username}>
            <TwitterAccountLink id={a.id} username={a.twitter_username} />
            <DropdownButton>
              <DropdownAction
                onClick={e => {
                  e.preventDefault();
                  this.props.unlinkTwitterAccount(a.id)
                }}
                confirmation="You can not revert this action. All campaign data will be lost. Are you sure?"
              >
                  Unlink
              </DropdownAction>
            </DropdownButton>
          </li>
        );
      };

      if (accounts.length == 0) {
        return <CleanDashboard />;
      } else {
        return (
          <div className="accounts">
            <h1>My accounts</h1>
            <ul className="list-unstyled">
              {accounts}
            </ul>
          </div>
        );
      }
    }
  }
}

const mapStateToProps = (state) => ({
  isFetching: state.twitterAccounts.isFetching,
  accounts: state.twitterAccounts.items
});
const mapDispatchToProps = {
  fetchTwitterAccounts,
  unlinkTwitterAccount
};

export default connect(mapStateToProps, mapDispatchToProps)(Accounts);
