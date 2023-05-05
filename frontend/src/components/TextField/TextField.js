import React from 'react'
import css from './TextField.less'

class TextField extends React.Component {
  constructor(props) {
    super(props);
  }

  hasError() {
    return this.props.showError && this.props.errorText != "";
  }

  render() {
    return (
      <div className="form-group">
        <label htmlFor={this.props.name}>{this.props.label}</label>
        <input onChange={this.props.onChange} value={this.props.value} type={this.props.type} name={this.props.name} className="form-control" id={this.props.name} />
        { this.hasError() ? <span className="help-block error">{this.props.errorText}</span> : null }
      </div>
    );
  }
}

export default TextField;
