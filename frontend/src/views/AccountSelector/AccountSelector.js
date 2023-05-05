import React from 'react';
import { Link } from 'react-router-dom';

const AccountSelector = ({accounts}) => {
  let items = accounts.map((a) => (
    <li key={a.twitter_username}>
      <Link to={`/dashboard/${a.twitter_username}`}>@{a.twitter_username}</Link>
    </li>
  ));

  return (
    <div>
      <p>Please select Twitter account:</p>
      <ul className="list-unstyled">
        {items}
      </ul>
    </div>
  );
}

export default AccountSelector;
