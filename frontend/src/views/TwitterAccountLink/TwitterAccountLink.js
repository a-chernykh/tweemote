import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';

import css from './TwitterAccountLink.less';

const TwitterAccountLink = ({ id, username }) => {
  return (
    <span className="twitter-account-link">
      <Link to={`/dashboard/${id}`}>@{username}</Link>
    </span>
  );
};

TwitterAccountLink.propTypes = {
  id: PropTypes.number.isRequired,
  username: PropTypes.string.isRequired
};

export default TwitterAccountLink;
