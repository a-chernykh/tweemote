// https://spin.atomicobject.com/2016/10/05/form-validation-react/

const ErrorMessages = {
  isRequired: (fieldName) => `${fieldName} is required`,
  shouldMatch: (otherFieldName) => {
    return (fieldName) => `${fieldName} must match ${otherFieldName}`;
  }
};

export const addValidation = (field, name, ...validations) => {
  return (state) => {
    for (let v of validations) {
      let errorMessageFunc = v(state[field], state);
      if (errorMessageFunc) {
        return {[field]: errorMessageFunc(name)};
      }
    }
  };
};

export const validate = (state, validations) => {
  return validations.reduce((errors, v) => {
    return Object.assign(errors, v(state));
  }, {});
};

export const required = (text) => {
  if (text) {
    return null;
  } else {
    return ErrorMessages.isRequired;
  }
};

export const matches = (fieldToMatch, fieldNameToMatch) => {
  return (text, state) => {
    if (text == state[fieldToMatch]) {
      return null;
    } else {
      return ErrorMessages.shouldMatch(fieldNameToMatch);
    }
  };
};
