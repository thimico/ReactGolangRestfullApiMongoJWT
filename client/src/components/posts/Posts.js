import React, { Fragment, useEffect } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import Spinner from '../common/Spinner';
import PostItem from './PostItem';
import PostForm from './PostForm';
import { getPosts } from '../../actions/postActions';

let loading = true
const Posts = ({ getPosts, 
    posts
}) => {
    loading = false
  useEffect(() => {
    getPosts();
  }, [getPosts]);

  return (
    <Fragment>
      {posts === null || loading ? (
        <Spinner />
      ) : (
    <Fragment>
      <h1 className='large text-primary'>Posts</h1>
      <p className='lead'>
        <i className='fas fa-user' /> Welcome to the community
      </p>
      <PostForm />
      <div className='posts'>
      {!!posts && posts.length > 0 ? 
               posts.map(post => (
                <PostItem key={post._id} post={post} />
              ))
                : (
        <h4>No Posts yet</h4>
        )}
        
      </div>
    </Fragment>
    
    )}
  </Fragment>
  );
};

Posts.propTypes = {
  getPosts: PropTypes.func.isRequired,
  post: PropTypes.object
};

const mapStateToProps = state => ({
  post: state.post
});

export default connect(
  mapStateToProps,
  { getPosts }
)(Posts);