import React from 'react'
import PropTypes from 'prop-types'
import css from './Signup.less'

import { Redirect } from 'react-router-dom';
import { Squares } from 'react-activity';

import update from 'immutability-helper';
import { connect } from 'react-redux'
import { addValidation, validate, required, matches } from 'lib/validation'
import { signupUser, signinUser } from 'state/actions/auth';
import SignupForm from 'views/SignupForm'

const validations = [
  addValidation("email", "E-mail", required),
  addValidation("password", "Password", required),
  addValidation("password_confirmation", "Password Confirmation", required, matches("password", "Password"))
]

class Signup extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      email: '',
      password: '',
      password_confirmation: '',
      showErrors: false
    };

    this.state.errors = validate(this.state, validations);

    this.signup = this.signup.bind(this);
    this.handleInputChange = this.handleInputChange.bind(this);
  }

  signup(e) {
    e.preventDefault();
    this.setState({showErrors: true});
    if (Object.keys(this.state.errors).length != 0) {
      return null;
    }

    this.props.signupUser(this.state.email,
                          this.state.password,
                          this.state.password_confirmation);
  }

  handleInputChange(e) {
    let newState = update(this.state, {
      [e.target.name]: {$set: e.target.value}
    });
    newState.errors = validate(newState, validations);
    this.setState(newState);
  }

  render() {
    if (this.props.isAuthenticated) {
      return (
        <Redirect to="/dashboard"/>
      );
    } else if (this.props.isRegistered) {
      this.props.signinUser(this.state.email, this.state.password);
      return <Squares size={40} />;
    } else {
      return (
        <SignupForm
          handleSubmit={this.signup}
          handleInputChange={this.handleInputChange}
          isRegistering={this.props.isRegistering}
          errors={this.state.errors}
          showErrors={this.state.showErrors}
        />
      );
    }
  }
};
Signup.contextTypes = {
  store: PropTypes.object
};

const mapStateToProps = (state) => ({
  isRegistering: state.session.isRegistering,
  isRegistered: state.session.isRegistered,
  isAuthenticated: state.session.isAuthenticated
});

const mapDispatchToProps = {
  signupUser,
  signinUser
};

export default connect(mapStateToProps, mapDispatchToProps)(Signup);
