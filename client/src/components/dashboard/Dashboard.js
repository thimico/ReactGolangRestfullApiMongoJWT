import React, { Component } from 'react'
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import Spinner from '../common/Spinner';
import { getCurrentProfile } from '../../actions/profileActions';
import Experience from './Experience';
import Education from './Education';

class Dashboard extends Component {
    componentDidMount() {
        this.props.getCurrentProfile();
    }

    render() {
        const { user } = this.props.auth;
        const { profile, loading } = this.props.profile;

        let dashboardContent;
        if (profile === null || loading) {
            dashboardContent = <Spinner />;
        } else {
            // Check if logged in user has profile data
            if (Object.keys(profile).length > 0) {
                dashboardContent = (
                <div>
                    <h1 className="display-4">Dashboard</h1>
                        <p className="lead text-muted">Welcome John Doe</p>
                        <div className="btn-group mb-4" role="group">
                            <Link to="/edit-profile" className="btn btn-light">
                            <i className="fas fa-user-circle text-info mr-1"></i> Edit Profile</Link>
                            <Link to="/add-experience" className="btn btn-light">
                            <i className="fab fa-black-tie text-info mr-1"></i>
                            Add Experience</Link>
                            <Link to="/add-education" className="btn btn-light">
                            <i className="fas fa-graduation-cap text-info mr-1"></i>
                            Add Education</Link>
                        </div>
                        if (!!profile.experience) {
                        <Experience
                            experience = {profile.experience}
                            />
                        }
                        if (!!profile.education) {
                            <Education education={profile.education} />
                        }
                        <div tyle={{ marginBottom: '60px' }}>
                            <button className="btn btn-danger">
                            Delete My Account
                            </button>
                        </div>
                    </div>
                )
            } else {
                 // User is logged in but has no profile
                dashboardContent = (
                    <div>
                        <p className="lead text-muted">Welcome {user.name}</p>
                        <p>You have not yet setup a profile, please add some info</p>
                        <Link to="/create-profile" className="btn btn-lg btn-info">
                        Create Profile
                        </Link>
                    </div>
                )
            }
        }


        return (
            <div className="dashboard">
                <div className="container">
                <div className="row">
                    <div className="col-md-12">
                        {dashboardContent}
                    </div>
                </div>
                </div>
            </div>
        )
    }
}

Dashboard.propTypes = {
    getCurrentProfile: PropTypes.func.isRequired,
    auth: PropTypes.object.isRequired,
    profile: PropTypes.object.isRequired
  };
  
  const mapStateToProps = state => ({
    profile: state.profile,
    auth: state.auth
  });

export default connect(mapStateToProps, { getCurrentProfile })(Dashboard)