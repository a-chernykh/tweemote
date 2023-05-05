import React from 'react';
import { Link } from 'react-router-dom';

const CleanDashboard = ({}) => {
  return (
    <div>
      <p>You don't have any Twitter accounts linked.</p>
      <p><Link className="btn btn-primary" to="/link/twitter">Link Twitter account</Link></p>
    </div>
  );
}

export default CleanDashboard;
