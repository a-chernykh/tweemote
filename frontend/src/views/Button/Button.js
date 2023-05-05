import React, { Component } from 'react';
import { render } from 'react-dom';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';

import { randomDomId } from 'lib/browser';
import ModalDialog from 'views/ModalDialog';

class Button extends Component {
  render() {
    let { onClick, confirmation, children, ...other } = this.props;

    if (confirmation) {
      let id = randomDomId();

      return (
        <div>
          <ModalDialog id={id} header="Confirmation" onConfirm={onClick}>{confirmation}</ModalDialog>
          <button data-toggle="modal" data-target={`#${id}`} type="button" {...other}>
            {children}
          </button>
        </div>
      );
    } else {
      return (
        <button onClick={onClick} type="button" {...other}>
          {children}
        </button>
      );
    }
  }
};

Button.propTypes = {
  children: PropTypes.any.isRequired,
  onClick: PropTypes.func.isRequired
};

export default Button;
