// https://blog.komand.com/how-to-render-components-outside-the-main-react-app

import React from 'react';
import { render } from 'react-dom';

const appendedElements = {};
const appendElementContainer = document.querySelector('.append-element-container');

function getAppendedElements() {
  const elements = [];

  const keys = Object.keys(appendedElements);
  const length = keys.length;

  if (length > 0) {
    keys.forEach((key) => {
      elements.push(React.cloneElement(appendedElements[key], { key: key }));
    });
  }

  return elements;
}

class AppendBodyComponent extends React.Component {
  constructor(props) {
    super(props);

    this.appendElementContainer = appendElementContainer;
  }

  setAppendElementId(id) {
    this.appendElementId = id;
  }

  updateAppendElement(content) {
    appendedElements[this.appendElementId] = content;

    this.updateAppendElements();
  }

  updateAppendElements() {
    render(<span>{getAppendedElements()}</span>, appendElementContainer);
  }

  removeAppendElement() {
    delete appendedElements[this.appendElementId];

    this.updateAppendElements();
  }

  componentDidMount() {
    this.updateSelf();
  }

  componentDidUpdate() {
    this.updateSelf();
  }

  componentWillUnmount() {
    this.removeAppendElement();
  }

  updateSelf() {
    this.updateAppendElement(this.bodyElements());
  }

  render() {
    // NOTE: since this is an append body component, we need to manage the rendering ourselves
    return null;
  }
};

export default AppendBodyComponent;
