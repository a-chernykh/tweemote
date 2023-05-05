import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';

import css from './Userbar.less';
import { signoutUser } from 'state/actions/auth';
import { fetchUserInfo } from 'state/actions/user';

class Userbar extends Component {
  componentDidMount() {
    this.props.fetchUserInfo();
  }

  render() {
    return (
      <ul className="nav navbar-nav navbar-user">
        <li className="dropdown">
          <a href="#" className="dropdown-toggle" data-toggle="dropdown" role="button">
            {this.props.email}{" "}
            <span className="caret"></span>
          </a>
          <ul className="dropdown-menu">
            <li><Link to="/accounts">My accounts</Link></li>
            <li><a onClick={this.props.signoutUser} href="#">Sign out</a></li>
          </ul>
        </li>
      </ul>
    );
  }
};

const mapStateToProps = (state) => ({
  email: state.profile.email
});
const actionsMap = {
  signoutUser,
  fetchUserInfo
};

export default connect(mapStateToProps, actionsMap)(Userbar);
