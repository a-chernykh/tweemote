import React, { Component } from 'react';
import { render } from 'react-dom';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';

import { randomDomId } from 'lib/browser';
import ModalDialog from 'views/ModalDialog';

class DropdownAction extends Component {
  render() {
    if (this.props.confirmation) {
      let id = randomDomId();

      return (
        <li>
          <ModalDialog id={id} header="Confirmation" onConfirm={this.props.onClick}>{this.props.confirmation}</ModalDialog>
          <a data-toggle="modal" data-target={`#${id}`} href="#">{this.props.children}</a>
        </li>
      );
    } else {
      return (
        <li>
          <a href="#" onClick={this.props.onClick}>{this.props.children}</a>
        </li>
      );
    }
  }
};

DropdownAction.propTypes = {
  children: PropTypes.any.isRequired,
  onClick: PropTypes.func.isRequired,
  confirmation: PropTypes.string
};

export default DropdownAction;
