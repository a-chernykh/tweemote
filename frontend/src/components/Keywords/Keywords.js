import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
import { connect } from 'react-redux';
import { Squares } from 'react-activity';
import pluralize from 'pluralize';

import css from './Keywords.less';
import Button from 'views/Button';
import AddKeyword from 'components/AddKeyword';
import { fetchKeywords, addKeyword, deleteKeyword } from 'state/actions/keywords';

class Keywords extends Component {
  constructor(props) {
    super(props);
    this.handleAdd = this.handleAdd.bind(this);
    this.handleDelete = this.handleDelete.bind(this);
  }

  componentWillMount() {
    this.props.fetchKeywords(this.props.campaign.id);
  }

  getKeywords() {
    return Object.values(this.props.keywords).filter(kw => kw.campaign_id.toString() == this.props.campaign.id).sort((a, b) => {
      let akw = a.keyword.toUpperCase();
      let bkw = b.keyword.toUpperCase();
      if (akw < bkw) {
        return -1;
      }
      if (akw > bkw) {
        return 1;
      }
      return 0;
    });
  }

  handleAdd(keyword) {
    this.props.addKeyword(this.props.campaign.id, keyword);
  }

  handleDelete(id) {
    this.props.deleteKeyword(this.props.campaign.id, id);
  }

  render() {
    let el = <Squares size={40} />;
    let kws = this.getKeywords();

    if (this.props.isFetched) {
      if (kws.length > 0) {
        let items = kws.map((kw) =>
          <tr key={kw.id} className={kw.isNew ? "success" : ""}>
            <td>
              <span className="label label-default">{kw.keyword}</span>
            </td>
            <td>
              {kw.impressions_count}
            </td>
            <td>
              {kw.followers_count}
            </td>
            <td>
              {kw.isDeleting ? (
                  <Squares size={20} />
                ) : (
                  <Button onClick={() => this.handleDelete(kw.id)} confirmation={["Are you sure that you want to delete keyword ", <strong key={kw.id}>{kw.keyword}</strong>, "?"]} type="button" className="btn btn-sm btn-danger">Delete</Button>
                )}
            </td>
          </tr>
        );
        el = (
          <table className="table table-hover">
            <thead>
              <tr>
                <th>Keyword</th>
                <th>Impressions</th>
                <th>Followers</th>
                <th>Actions</th>
              </tr>
            </thead>
            <tfoot>
              <tr>
                <td colSpan="4">
                  {pluralize('keyword', kws.length, true)}
                </td>
              </tr>
            </tfoot>
            <tbody>
              {items}
            </tbody>
          </table>
        );
      } else {
        el = <p>You don't have any keywords defined.</p>;
      }
    }

    return <div className="keywords">
             <h3>Keywords</h3>
             <AddKeyword disable={this.props.isAdding} onAdd={this.handleAdd} />
             {el}
           </div>;
  }
}

const mapStateToProps = (state) => ({
  keywords:  state.keywords.items,
  isFetched: state.keywords.isFetched,
  isAdding:  state.keywords.isAdding
});
const mapDispatchToProps = {
  fetchKeywords,
  addKeyword,
  deleteKeyword
};

export default connect(mapStateToProps, mapDispatchToProps)(Keywords);
