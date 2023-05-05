import React from 'react';
import { Link } from 'react-router-dom';

import UserDashboard from 'components/UserDashboard';

const Dashboard = ({ children }) => {
  return (
    <div>
      <UserDashboard />
    </div>
  );
}

export default Dashboard;
