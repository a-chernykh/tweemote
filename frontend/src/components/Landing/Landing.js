import React from 'react';
import { Link } from 'react-router-dom'

import Header from '../Header'

import TwitterSmall from './twitter-small.png'

class Landing extends React.Component {
  render() {
    return (
      <div className="landing">
        <section className="header">
          <div className="container">
            <div className="row">
              <div className="col-lg-12">
                <img className="hidden-xs hidden-sm pull-left" src={TwitterSmall} alt="Get Twitter followers" />

                <div className="headline">
                  <h1>Get Twitter followers with precise targeting</h1>
                  <h2>Reactive Boost interacts with tweets and users which are relevant to your product. You get more followers because engaged users are more likely to be interested in your account.</h2>
                  <p>
                    <Link to='/signup' className="btn btn-primary btn-lg">Sign Up for free trial</Link>
                  </p>
                  <p>Already have an account? <a href="#">Sign in</a>.</p>
                </div>
              </div>
            </div>
          </div>
        </section>

        <section id="how" className="how">
          <div className="container">
            <div className="row">
              <div className="col-lg-12">
                <h1>How it works</h1>
                <dl>
                  <dt>1. Configure keywords and hashtags</dt>
                  <dd>Use our keyword selection tool to setup keywords which are relevant to your product. Reactive Boost will monitor and interact with matching tweets.</dd>

                  <dt>2. Select relevant Twitter accounts</dt>
                  <dd>Using powerful rule engine you can configure keywords to find Twitter users you are interested in.</dd>

                  <dt>3. Monitor your growth</dt>
                  <dd>Powerful Analytics engine provides insights about how campaign performs.</dd>
                </dl>
              </div>
            </div>
          </div>
        </section>

        <section id="features" className="features">
          <div className="container">
            <div className="row">
              <div className="col-lg-12">
                <h1>Features</h1>

                <div className="row features-list">
                  <div className="col-lg-4 feature">
                    <i className="fa fa-fw fa-binoculars" aria-hidden="true"></i>
                    <div className="feature-content">
                      <h2>Keywords research</h2>
                      <p>We will help to come up with relevant keywords by showing matching tweets and users in real time.</p>
                    </div>
                  </div>
                  <div className="col-lg-4 feature">
                    <i className="fa fa-fw fa-users" aria-hidden="true"></i>
                    <div className="feature-content">
                      <h2>Multiple accounts</h2>
                      <p>You can use multiple Twitter accounts (depending on plan).</p>
                    </div>
                  </div>
                  <div className="col-lg-4 feature">
                    <i className="fa fa-fw fa-line-chart" aria-hidden="true"></i>
                    <div className="feature-content">
                      <h2>Analytics</h2>
                      <p>View detailed stats on how your campaign performs. Day to day stats on number of followers and impressions.</p>
                    </div>
                  </div>
                </div>
                <div className="row features-list">
                  <div className="col-lg-4 feature">
                    <i className="fa fa-fw fa-stop" aria-hidden="true"></i>
                    <div className="feature-content">
                      <h2>Stop words</h2>
                      <p>Configure stop words for better targeting. Explicit content is ignored by default.</p>
                    </div>
                  </div>
                  <div className="col-lg-4 feature">
                    <i className="fa fa-fw fa-graduation-cap" aria-hidden="true"></i>
                    <div className="feature-content">
                      <h2>Machine learning</h2>
                      <p>Reactive Boost learns which interactions leaded to conversion and will apply this knowledge for future engagements.</p>
                    </div>
                  </div>
                  <div className="col-lg-4 feature">
                    <i className="fa fa-fw fa-undo" aria-hidden="true"></i>
                    <div className="feature-content">
                      <h2>Automatic unfollowing</h2>
                      <p>If user did not follow you in a month, you can opt in for automatic unfollowing.</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>
      </div>
    );
  }
}

export default Landing;
