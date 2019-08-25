import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Link, withRouter } from 'react-router-dom';
import PropTypes from 'prop-types';
import { createProfile } from '../../actions/profileActions';

class CreateProfile extends Component {
  constructor(props) {
    super(props);
    this.state = {
        displaySocialInputs: false,
        handle: '',
        company: "",
        website: "",
        location: "",
        status: "",
        skills: "",
        githubusername: "",
        bio: "",
        twitter: "",
        facebook: "",
        linkedin: "",
        youtube: "",
        instagram: "",
        errors: {}
    };

    this.onChange = this.onChange.bind(this);
    this.onSubmit = this.onSubmit.bind(this);
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.errors) {
      this.setState({ errors: nextProps.errors });
    }
  }

  onSubmit(e) {
    e.preventDefault();

    const data = {

        handle: this.state.handle,
        company: this.state.company,
        website: this.state.website,
        location: this.state.location,
        status: this.state.status,
        skills: this.state.skills,
        githubusername: this.state.githubusername,
        bio: this.state.bio,
        twitter: this.state.twitter,
        facebook: this.state.facebook,
        linkedin: this.state.linkedin,
        youtube: this.state.youtube,
        instagram: this.state.instagram
    };

    this.props.createProfile(data, this.props.history);
  }

  onChange(e) {
    this.setState({ [e.target.name]: e.target.value });
  }

  render() {
    const { errors, displaySocialInputs } = this.state;

    let socialInputs;

    if (displaySocialInputs) {
        socialInputs = (
            <div>
                <div className="input-group mb-3">
                    <div className="input-group-prepend">
                            <span className="input-group-text">
                            <i className="fab fa-twitter"></i>
                            </span>
                    </div>
                    <input type="text" className="form-control form-control-lg" placeholder="Twitter Profile URL" name="twitter"
                           value={this.state.twitter}
                           onChange={this.onChange}/>
                </div>

                <div className="input-group mb-3">
                    <div className="input-group-prepend">
                            <span className="input-group-text">
                            <i className="fab fa-facebook"></i>
                            </span>
                    </div>
                    <input type="text" className="form-control form-control-lg" placeholder="Facebook Page URL" name="facebook"
                           value={this.state.facebook}
                           onChange={this.onChange}/>
                </div>

                <div className="input-group mb-3">
                    <div className="input-group-prepend">
                            <span className="input-group-text">
                            <i className="fab fa-linkedin"></i>
                            </span>
                    </div>
                    <input type="text" className="form-control form-control-lg" placeholder="Linkedin Profile URL" name="linkedin"
                           value={this.state.linkedin}
                           onChange={this.onChange}/>
                </div>

                <div className="input-group mb-3">
                    <div className="input-group-prepend">
                            <span className="input-group-text">
                            <i className="fab fa-youtube"></i>
                            </span>
                    </div>
                    <input type="text" className="form-control form-control-lg" placeholder="YouTube Channel URL" name="youtube"
                           value={this.state.youtube}
                           onChange={this.onChange}/>
                </div>

                <div className="input-group mb-3">
                    <div className="input-group-prepend">
                            <span className="input-group-text">
                            <i className="fab fa-instagram"></i>
                            </span>
                    </div>
                    <input type="text" className="form-control form-control-lg" placeholder="Instagram Page URL" name="instagram"
                           value={this.state.instagram}
                           onChange={this.onChange}/>
                </div>
            </div>
        )
    }

    return (
        <div className="create-profile">
            <div className="container">
                <div className="row">
                    <div className="col-md-8 m-auto">
                        <Link to="/dashboard" className="btn btn-light">
                            Go Back
                        </Link>
                        <h1 className="display-4 text-center">Create Your Profile</h1>
                        <p className="lead text-center">Let's get some information to make your profile stand out</p>
                        <small className="d-block pb-3">* = required field</small>
                        <form className='form' onSubmit={this.onSubmit}>
                            <div className="form-group">
                                <input type="text" className="form-control form-control-lg" placeholder="* Profile handle" name="handle" required
                                       value={this.state.handle} onChange={this.onChange}/>
                                <small className="form-text text-muted">A unique handle for your profile URL. Your full name, company name, nickname, etc (This CAN'T be changed later)</small>
                            </div>
                            <div className="form-group">
                                <select className="form-control form-control-lg" name='status' value={this.state.status} onChange={this.onChange}>
                                    <option value='0'>* Select Professional Status</option>
                                    <option value='Developer'>Developer</option>
                                    <option value='Junior Developer'>Junior Developer</option>
                                    <option value='Senior Developer'>Senior Developer</option>
                                    <option value='Manager'>Manager</option>
                                    <option value='Student or Learning'>Student or Learning</option>
                                    <option value='Instructor'>Instructor or Teacher</option>
                                    <option value='Intern'>Intern</option>
                                    <option value='Other'>Other</option>
                                </select>
                                <small className="form-text text-muted">Give us an idea of where you are at in your career</small>
                                {errors.status && (
                                    <div className="invalid-feedback">{errors.status}</div>
                                )}
                            </div>
                            <div className="form-group">
                                <input type="text" className="form-control form-control-lg" placeholder="Company" name="company"
                                       value={this.state.company}
                                       onChange={this.onChange}/>
                                <small className="form-text text-muted">Could be your own company or one you work for</small>
                            </div>
                            <div className="form-group">
                                <input type="text" className="form-control form-control-lg" placeholder="Website" name="website"
                                       value={this.state.website}
                                       onChange={this.onChange}/>
                                <small className="form-text text-muted">Could be your own or a company website</small>
                            </div>
                            <div className="form-group">
                                <input type="text" className="form-control form-control-lg" placeholder="Location" name="location"
                                       value={this.state.location}
                                       onChange={this.onChange}/>
                                <small className="form-text text-muted">City & state suggested (eg. Boston, MA)</small>
                            </div>
                            <div className="form-group">
                                <input type="text" className="form-control form-control-lg" placeholder="Skills" name="skills"
                                       value={this.state.skills}
                                       onChange={this.onChange}/>
                                <small className="form-text text-muted">Please use comma separated values (eg. HTML,CSS,JavaScript,PHP)</small>
                            </div>
                            <div className="form-group">
                                <input type="text" className="form-control form-control-lg" placeholder="Github Username" name="githubusername"
                                       value={this.state.githubusername}
                                       onChange={this.onChange}/>
                                <small className="form-text text-muted">If you want your latest repos and a Github link, include your username</small>
                            </div>
                            <div className="form-group">
                                <textarea className="form-control form-control-lg" placeholder="A short bio of yourself" name="bio"
                                          value={this.state.bio}
                                          onChange={this.onChange}></textarea>
                                <small className="form-text text-muted">Tell us a little about yourself</small>
                            </div>

                            <div className="mb-3">
                                <button type="button"
                                        onClick={() => {
                                            this.setState(prevState => ({
                                                displaySocialInputs: !prevState.displaySocialInputs
                                            }));
                                        }}
                                        className="btn btn-light">Add Social Network Links</button>
                                <span className="text-muted">Optional</span>
                            </div>
                            {socialInputs}

                            <input type="submit"  value="Submit" className="btn btn-info btn-block mt-4" />
                        </form>
                    </div>
                </div>
            </div>
        </div>
    );
  }
}
CreateProfile.propTypes = {
    profile: PropTypes.object.isRequired,
    errors: PropTypes.object.isRequired
};
const mapStateToProps = state => ({
    profile: state.profile,
    errors: state.errors
});
export default connect(mapStateToProps, { createProfile })(
  withRouter(CreateProfile)
);