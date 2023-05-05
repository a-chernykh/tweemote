import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
import { connect } from 'react-redux';
import { Squares } from 'react-activity';
import PropTypes from 'prop-types';

import css from './AddKeyword.less';

class AddKeyword extends Component {
  constructor(props) {
    super(props);
    this.state = { keyword: '' };
    this.handleChange = this.handleChange.bind(this);
    this.handleAdd = this.handleAdd.bind(this);
  }

  handleChange(e) {
    e.preventDefault();
    this.setState({
      keyword: e.target.value
    });
  }

  handleAdd(e) {
    e.preventDefault();
    if (this.props.onAdd) {
      this.props.onAdd(this.state.keyword);
      this.setState({
        keyword: ""
      });
      document.getElementById("new-keyword").focus();
    }
  }

  render() {
    return (
      <form className="form-horizontal add-keyword">
        <div className="form-group">
          <div className="col-xs-8">
            <input value={this.state.keyword}
                   onChange={this.handleChange}
                   type="text"
                   name="keyword"
                   id="new-keyword"
                   className="form-control"
                   placeholder="Keyword" />
          </div>
          <div className="col-xs-4">
            <button disabled={this.props.disable} type="submit" className="btn btn-primary" onClick={this.handleAdd}>
              {this.props.disable ? "Please wait..." : "Add keyword" }
            </button>
          </div>
        </div>
      </form>
    );
  }
}

AddKeyword.propTypes = {
  onAdd: PropTypes.func,
  disable: PropTypes.bool
};

const mapDispatchToProps = {
};

export default connect(null, mapDispatchToProps)(AddKeyword);
