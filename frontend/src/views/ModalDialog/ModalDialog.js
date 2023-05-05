import React from 'react';
import AppendBodyComponent from 'components/AppendBody';
import PropTypes from 'prop-types';
import css from './ModalDialog.less';

class ModalDialog extends AppendBodyComponent {
  constructor(props) {
    super(props);
    this.setAppendElementId(props.id);
  }

  bodyElements() {
    return (
      <div className="modal fade" id={this.props.id} tabIndex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
        <div className="modal-dialog">
          <div className="modal-content">
            <div className="modal-header">
              {this.props.header}
            </div>
            <div className="modal-body">
              {this.props.children}
            </div>
            <div className="modal-footer">
              <button type="button" className="btn btn-default" data-dismiss="modal">Cancel</button>
              <a className="btn btn-danger btn-ok" data-dismiss="modal" onClick={this.props.onConfirm}>Delete</a>
            </div>
          </div>
        </div>
      </div>
    );
  }
};

ModalDialog.propTypes = {
  children: PropTypes.any.isRequired,
  header: PropTypes.string.isRequired,
  onConfirm: PropTypes.func
};

export default ModalDialog;
