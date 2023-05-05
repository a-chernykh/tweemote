import React from 'react';
import PropTypes from 'prop-types';

import { Link } from 'react-router-dom'
import TextField from 'components/TextField'
import SubmitButton from 'components/SubmitButton'

const SignupForm = ({handleSubmit,
                     handleInputChange,
                     isRegistering,
                     errors,
                     showErrors}) => {
  return (
    <form onSubmit={handleSubmit}>
      <h1>Sign Up</h1>

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
      <TextField
        type="password"
        name="password_confirmation"
        label="Password Confirmation"
        onChange={handleInputChange}
        errorText={errors.password_confirmation}
        showError={showErrors}
      />

      <SubmitButton
        disable={isRegistering}
        label="Sign Up"
      />

      <p>
        Already have an account? <Link to="/signin">Sign in</Link>
      </p>
    </form>
  );
};

SignupForm.propTypes = {
  handleSubmit: PropTypes.func.isRequired,
  handleInputChange: PropTypes.func.isRequired,
  isRegistering: PropTypes.bool,
  errors: PropTypes.objectOf(PropTypes.string).isRequired,
  showErrors: PropTypes.bool.isRequired
};

export default SignupForm;
