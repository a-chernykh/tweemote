import React from 'react'
import { Squares } from 'react-activity';

import css from './SubmitButton.less'

class SubmitButton extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div>
        <button disabled={this.props.disable ? "disabled" : false} type="submit" className="btn btn-primary">{this.props.label}</button>
        {this.props.disable ? <Squares size={18} /> : null}
      </div>
    );
  }
}

export default SubmitButton;
