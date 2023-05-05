import React from 'react';
import PropTypes from 'prop-types';

import { Link } from 'react-router-dom';
import TextField from 'components/TextField';
import SubmitButton from 'components/SubmitButton';

const SigninForm = ({handleSubmit,
                     handleInputChange,
                     isAuthenticating,
                     errors,
                     showErrors}) => {
  return (
    <form onSubmit={handleSubmit}>
      <h1>Sign In</h1>

      <TextField
        type="text"
        name="email"
        label="E-mail"
        onChange={handleInputChange}
        errorText={errors.email}
        showError={showErrors}
      />
      <TextField
        type="password"
        name="password"
        label="Password"
        onChange={handleInputChange}
        errorText={errors.password}
        showError={showErrors}
      />

      <SubmitButton
        disable={isAuthenticating}
        label="Sign In"
      />

      <p>
        Don't have an account? <Link to="/signup">Sign up</Link>
      </p>
    </form>
  );
};
SigninForm.propTypes = {
  handleSubmit: PropTypes.func.isRequired,
  handleInputChange: PropTypes.func.isRequired,
  isAuthenticating: PropTypes.bool,
  errors: PropTypes.objectOf(PropTypes.string).isRequired,
  showErrors: PropTypes.bool.isRequired
};
export default SigninForm;
