import React, { Component } from 'react';
import Plotly from 'lib/plotly';
import { connect } from 'react-redux';
import { Squares } from 'react-activity';

import { fetchStats } from 'state/actions/stats';

class CampaignDashboard extends Component {
  componentWillMount() {
    this.rendered = false;
    this.props.fetchStats(this.props.campaign.id);
  }

  componentDidUpdate() {
    if (!this.rendered && document.getElementById("likes-graph")) {
      let stats = this.props.stats[this.props.campaign.id];

      let impressionsTrace = {
        x: stats.map(s => s.day),
        y: stats.map(s => s.impressions),
        type: 'bar'
      };
      let followersTrace = {
        x: stats.map(s => s.day),
        y: stats.map(s => s.followers),
        type: 'bar'
      };

      let impressionsLayout = {
        title: 'New likes',
        yaxis: {
          title: 'Number of likes',
          side: 'left'
        }
      };

      Plotly.newPlot('likes-graph', [impressionsTrace], impressionsLayout);

      let followersLayout = {
        title: 'New followers',
        yaxis: {
          title: 'Number of new followers',
          side: 'left'
        }
      };

      Plotly.newPlot('followers-graph', [followersTrace], followersLayout);

      this.rendered = true;
    }
  }

  render() {
    if (this.props.isFetching) {
      return <Squares size={40} />;
    } else {
      return (
        <div>
          <div id="likes-graph"></div>
          <div id="followers-graph"></div>
        </div>
      );
    }
  }
}

const mapStateToProps = (state) => ({
  stats:      state.stats.items,
  isFetching: state.stats.isFetching,
  isFetched:  state.stats.isFetched,
  isAdding:   state.stats.isAdding
});
const mapDispatchToProps = {
  fetchStats
};

export default connect(mapStateToProps, mapDispatchToProps)(CampaignDashboard);
