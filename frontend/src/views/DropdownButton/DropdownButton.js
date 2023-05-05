import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';

const DropdownButton = ({ name, children }) => {
  let title;

  if (name) {
    title = <span className="caret"></span>;
  } else {
    title = <span className="glyphicon glyphicon-menu-hamburger" aria-hidden="true"></span>;
  }

  return (
    <div className="btn-group">
      <button type="button" className="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
        {title}
      </button>
      <ul className="dropdown-menu">
        {children}
      </ul>
    </div>
  );
};

DropdownButton.propTypes = {
  children: PropTypes.element.isRequired
};

export default DropdownButton;
