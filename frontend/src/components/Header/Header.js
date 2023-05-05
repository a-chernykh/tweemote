import React from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';

import css from './Header.less';
import Navbar from 'views/Navbar';

class Header extends React.Component {
  render() {
    let brandLink;

    if (this.props.isAuthenticated) {
      brandLink = '/dashboard';
    } else {
      brandLink = '/';
    }

    return (
      <header>
        <div className="navbar navbar-default navbar-fixed-top" role="navigation">
          <div className="container-fluid">
            <div className="navbar-header">
              <button type="button" className="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
                <span className="sr-only">Toggle navigation</span>
                <span className="icon-bar"></span>
                <span className="icon-bar"></span>
                <span className="icon-bar"></span>
              </button>
              <Link className="navbar-brand" to={brandLink}>Reactive Boost</Link>
            </div>
            <div className="navbar-collapse collapse navbar-right">
              <Navbar />
            </div>
          </div>
        </div>
      </header>
    );
  }
};

const mapStateToProps = (state) => ({
  isAuthenticated: state.session.isAuthenticated
});

export default connect(mapStateToProps)(Header);
