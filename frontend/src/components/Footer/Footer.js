import React from 'react';
import PropTypes from 'prop-types';

import css from './Footer.less'

class Footer extends React.Component {
  render() {
    return (
      <footer id="contact">
        <div className="container">
          <div className="row">
            <div className="col-lg-12">
              <p className="pull-left">Copyright &copy; 2017 Reactive Boost. All rights reserved.</p>
              <p className="pull-right">Questions? Contact us at <a href="mailto:andrey@reactiveboost.com">andrey@reactiveboost.com</a></p>
            </div>
          </div>
        </div>
      </footer>
    );
  }
};

export default Footer;
