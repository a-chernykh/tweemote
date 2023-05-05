import React from 'react';
import PropTypes from 'prop-types';
import {
  BrowserRouter as Router,
  Route,
  Redirect
} from 'react-router-dom';
import NotificationSystem from 'react-notification-system';
import { Provider } from 'react-redux';

import { notifier } from 'lib/notifier';
import store from 'state/store';
import PrivateRoute from 'components/PrivateRoute';

import Header from 'components/Header';
import Footer from 'components/Footer';
import Landing from 'components/Landing';
import Signup from 'components/Signup';
import Signin from 'components/Signin';
import Dashboard from 'views/Dashboard';
import LinkTwitter from 'components/LinkTwitter';
import LinkTwitterCallback from 'components/LinkTwitterCallback';
import Accounts from 'components/Accounts';
import CampaignSelector from 'components/CampaignSelector';
import CampaignSection from 'views/CampaignSection';

import bootstrapJs from 'bootstrap/dist/js/bootstrap';
import css from './Layout.less';

class Layout extends React.Component {
  componentDidMount() {
    notifier.setAdapter(this.refs.notificationSystem);
  }

  render() {
    return (
      <Provider store={store}>
        <Router>
          <div id="all">
            <NotificationSystem ref="notificationSystem" />

            <Header />

            <Route exact path="/" component={Landing} />

            <div className="container-fluid">
              <div className="row">
                <div className="col-lg-12">

                  <Route path="/signup" component={Signup} />
                  <Route path="/signin" component={Signin} />

                  <PrivateRoute exact path="/dashboard" component={Dashboard} />
                  <PrivateRoute exact path="/link/twitter" component={LinkTwitter} />
                  <PrivateRoute exact path="/link/twitter/callback" component={LinkTwitterCallback} />
                  <PrivateRoute path="/accounts" component={Accounts} />

                  <PrivateRoute exact path="/dashboard/:id" component={CampaignSelector} />
                  <PrivateRoute path="/dashboard/:id/:campaign_id" component={CampaignSection} />

                </div>
              </div>
            </div>

            <Footer />
          </div>
        </Router>
      </Provider>
    );
  }
}

export default Layout;
