import React from 'react';
import { Redirect } from 'react-router-dom';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import update from 'immutability-helper';

import css from './Signin.less';
import { addValidation, validate, required, matches } from 'lib/validation';
import SigninForm from 'views/SigninForm';
import { signinUser } from 'state/actions/auth';

const validations = [
  addValidation("email", "E-mail", required),
  addValidation("password", "Password", required),
]

class Signin extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      email: '',
      password: '',
      showErrors: false
    };

    this.state.errors = validate(this.state, validations);

    this.Signin = this.Signin.bind(this);
    this.handleInputChange = this.handleInputChange.bind(this);
  }

  Signin(e) {
    e.preventDefault();

    this.setState({showErrors: true});
    if (Object.keys(this.state.errors).length != 0) {
      return null;
    }

    this.props.onSignInClick(this.state.email, this.state.password);
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
      const { from } = this.props.location.state || { from: { pathname: '/dashboard' } }
      return (
        <Redirect to={from}/>
      );
    } else {
      return (
        <SigninForm
          handleSubmit={this.Signin}
          handleInputChange={this.handleInputChange}
          isAuthenticating={this.props.isAuthenticating}
          errors={this.state.errors}
          showErrors={this.state.showErrors}
        />
      );
    }
  }
};
Signin.contextTypes = {
  store: PropTypes.object
};

const mapStateToProps = (state) => ({
  isAuthenticating: state.session.isAuthenticating,
  isAuthenticated: state.session.isAuthenticated
});

export default connect(mapStateToProps, { onSignInClick: signinUser })(Signin);
