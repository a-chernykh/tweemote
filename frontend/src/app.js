import React from 'react';
import ReactDOM from 'react-dom';
import { AppContainer } from 'react-hot-loader';

import Layout from './components/Layout';

const render = (Component) => {
  ReactDOM.render(
    <AppContainer>
      <Component />
    </AppContainer>,
    document.getElementById('root')
  );
};

render(Layout);

if (module.hot) {
  module.hot.accept('./components/Layout', () => {
    const NextLayout = require('./components/Layout').default;
    ReactDOM.render(
      <AppContainer>
        <NextLayout />
      </AppContainer>,
      document.getElementById('root')
    );
  });
}
