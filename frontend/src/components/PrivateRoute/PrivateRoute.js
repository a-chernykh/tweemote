import React from 'react';
import {
  BrowserRouter as Router,
  Route,
  Redirect
} from 'react-router-dom';

import store from 'state/store';

const renderMergedProps = (component, ...rest) => {
  const finalProps = Object.assign({}, ...rest);
  return (
    React.createElement(component, finalProps)
  );
};

class PrivateRoute extends React.Component {
  constructor(props) {
    super(props);

    let { component: Component, ...rest } = props;

    this.Component = Component;
    this.rest = rest;
  }

  componentWillMount() {
    this.setState({
      isAuthenticated: store.getState().session.isAuthenticated
    });
  }

  componentDidMount() {
    const updateState = () => {
      this.setState({
        isAuthenticated: store.getState().session.isAuthenticated
      });
    };
    this.unsubscribe = store.subscribe(function() {
      updateState();
    });
  }

  componentWillUnmount() {
    if (this.unsubscribe) {
      this.unsubscribe();
    }
  }

  render() {
    return <Route {...this.rest} render={props => {
      if (store.getState().session.isAuthenticated) {
        return renderMergedProps(this.Component, props, this.rest);
      } else {
        return <Redirect to={{
          pathname: '/signin',
          state: { from: props.location }
        }} />;
      }
    }} />;
  }
};

export default PrivateRoute;
