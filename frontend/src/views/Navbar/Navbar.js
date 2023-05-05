import React from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';

import css from './Navbar.less';
import Userbar from 'views/Userbar';

const Navbar = ({ isAuthenticated, onSignOutClick }) => {
  if (isAuthenticated) {
    return (<Userbar />);
  } else {
    return (
      <ul className="nav navbar-nav navbar-guest">
        <li><a href="#home">Home</a></li>
        <li><a href="#how">How it works</a></li>
        <li><a href="#features">Features</a></li>
        <li><a href="#contact">Contact</a></li>
        <li><form><Link to="/signup" className="btn btn-primary navbar-btn">Sign Up</Link></form></li>
        <li><form><Link to="/signin" className="btn btn-default navbar-btn">Sign In</Link></form></li>
      </ul>
    );
  }
};

const mapStateToProps = (state) => ({
  isAuthenticated: state.session.isAuthenticated
});

export default connect(mapStateToProps)(Navbar);
