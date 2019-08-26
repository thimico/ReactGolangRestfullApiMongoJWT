import React, { Fragment, useState, useEffect } from 'react';
import { Link, withRouter } from 'react-router-dom';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { updateProfile, getCurrentProfile } from '../../actions/profileActions';

const EditProfile = ({
                         profile: { profile, loading },
                         updateProfile,
                         getCurrentProfile,
                         history
                     }) => {
    const [formData, setFormData] = useState({
        id: '',
        handle: '',
        company: '',
        website: '',
        location: '',
        status: '',
        skills: '',
        githubusername: '',
        bio: '',
        twitter: '',
        facebook: '',
        linkedin: '',
        youtube: '',
        instagram: ''
    });

    const [displaySocialInputs, toggleSocialInputs] = useState(false);

    useEffect(() => {
        getCurrentProfile();

        setFormData({
            id: loading || !profile.id ? '' : profile.id,
            handle: loading || !profile.handle ? '' : profile.handle,
            company: loading || !profile.company ? '' : profile.company,
            website: loading || !profile.website ? '' : profile.website,
            location: loading || !profile.location ? '' : profile.location,
            status: loading || !profile.status ? '' : profile.status,
            skills: loading || !profile.skills ? '' : profile.skills,
            githubusername:
                loading || !profile.githubusername ? '' : profile.githubusername,
            bio: loading || !profile.bio ? '' : profile.bio,
            twitter: loading || !profile.social ? '' : profile.social.twitter,
            facebook: loading || !profile.social ? '' : profile.social.facebook,
            linkedin: loading || !profile.social ? '' : profile.social.linkedin,
            youtube: loading || !profile.social ? '' : profile.social.youtube,
            instagram: loading || !profile.social ? '' : profile.social.instagram
        });
    }, [loading, getCurrentProfile]);

    const {
        id,
        handle,
        company,
        website,
        location,
        status,
        skills,
        githubusername,
        bio,
        twitter,
        facebook,
        linkedin,
        youtube,
        instagram
    } = formData;

    const onChange = e =>
        setFormData({ ...formData, [e.target.name]: e.target.value });

    const onSubmit = e => {
        e.preventDefault();
        updateProfile(formData, history, true);
    };

    return (
        <Fragment>
            <h1 className='large text-primary'>Edit Your Profile</h1>
            <p className='lead'>
                <i className='fas fa-user' /> Add some changes to your profile
            </p>
            <small>* = required field</small>
            <form className='form' onSubmit={e => onSubmit(e)}>
                <input type="hidden" name="id" value={id} onChange={e => onChange(e)}/>
                <div className="form-group">
                    <input type="text" className="form-control form-control-lg" placeholder="* Profile handle" name="handle" required
                            value={handle} onChange={e => onChange(e)}/>
                    <small className="form-text text-muted">A unique handle for your profile URL. Your full name, company name, nickname, etc (This CAN'T be changed later)</small>
                </div>
                <div className='form-group'>
                    <select name='status' className="form-control form-control-lg" value={status} onChange={e => onChange(e)}>
                        <option>* Select Professional Status</option>
                        <option value='Developer'>Developer</option>
                        <option value='Junior Developer'>Junior Developer</option>
                        <option value='Senior Developer'>Senior Developer</option>
                        <option value='Manager'>Manager</option>
                        <option value='Student or Learning'>Student or Learning</option>
                        <option value='Instructor'>Instructor or Teacher</option>
                        <option value='Intern'>Intern</option>
                        <option value='Other'>Other</option>
                    </select>
                    <small className='form-text'>
                        Give us an idea of where you are at in your career
                    </small>
                </div>
                <div className='form-group'>
                    <input className="form-control form-control-lg"
                        type='text'
                        placeholder='Company'
                        name='company'
                        value={company}
                        onChange={e => onChange(e)}
                    />
                    <small className='form-text'>
                        Could be your own company or one you work for
                    </small>
                </div>
                <div className='form-group'>
                    <input className="form-control form-control-lg"
                        type='text'
                        placeholder='Website'
                        name='website'
                        value={website}
                        onChange={e => onChange(e)}
                    />
                    <small className='form-text'>
                        Could be your own or a company website
                    </small>
                </div>
                <div className='form-group'>
                    <input className="form-control form-control-lg"
                        type='text'
                        placeholder='Location'
                        name='location'
                        value={location}
                        onChange={e => onChange(e)}
                    />
                    <small className='form-text'>
                        City & state suggested (eg. Boston, MA)
                    </small>
                </div>
                <div className='form-group'>
                    <input className="form-control form-control-lg"
                        type='text'
                        placeholder='* Skills'
                        name='skills'
                        value={skills}
                        onChange={e => onChange(e)}
                    />
                    <small className='form-text'>
                        Please use comma separated values (eg. HTML,CSS,JavaScript,PHP)
                    </small>
                </div>
                <div className='form-group'>
                    <input className="form-control form-control-lg"
                        type='text'
                        placeholder='Github Username'
                        name='githubusername'
                        value={githubusername}
                        onChange={e => onChange(e)}
                    />
                    <small className='form-text'>
                        If you want your latest repos and a Github link, include your
                        username
                    </small>
                </div>
                <div className='form-group'>
          <textarea className="form-control form-control-lg"
              placeholder='A short bio of yourself'
              name='bio'
              value={bio}
              onChange={e => onChange(e)}
          />
                    <small className='form-text'>Tell us a little about yourself</small>
                </div>

                <div className='my-2'>
                    <button className="btn btn-info btn-block mt-4"
                        onClick={() => toggleSocialInputs(!displaySocialInputs)}
                        type='button'
                        className='btn btn-light'
                    >
                        Add Social Network Links
                    </button>
                    <span>Optional</span>
                </div>

                {displaySocialInputs && (
                    <Fragment>
                        <div className='form-group social-input'>
                            <i className='fab fa-twitter fa-2x' />
                            <input
                                type='text'
                                placeholder='Twitter URL'
                                name='twitter'
                                value={twitter}
                                onChange={e => onChange(e)}
                            />
                        </div>

                        <div className='form-group social-input'>
                            <i className='fab fa-facebook fa-2x' />
                            <input
                                type='text'
                                placeholder='Facebook URL'
                                name='facebook'
                                value={facebook}
                                onChange={e => onChange(e)}
                            />
                        </div>

                        <div className='form-group social-input'>
                            <i className='fab fa-youtube fa-2x' />
                            <input
                                type='text'
                                placeholder='YouTube URL'
                                name='youtube'
                                value={youtube}
                                onChange={e => onChange(e)}
                            />
                        </div>

                        <div className='form-group social-input'>
                            <i className='fab fa-linkedin fa-2x' />
                            <input
                                type='text'
                                placeholder='Linkedin URL'
                                name='linkedin'
                                value={linkedin}
                                onChange={e => onChange(e)}
                            />
                        </div>

                        <div className='form-group social-input'>
                            <i className='fab fa-instagram fa-2x' />
                            <input
                                type='text'
                                placeholder='Instagram URL'
                                name='instagram'
                                value={instagram}
                                onChange={e => onChange(e)}
                            />
                        </div>
                    </Fragment>
                )}

                <input type='submit' className='btn btn-primary my-1' />
                <Link className='btn btn-light my-1' to='/dashboard'>
                    Go Back
                </Link>
            </form>
        </Fragment>
    );
};

EditProfile.propTypes = {
    updateProfile: PropTypes.func.isRequired,
    getCurrentProfile: PropTypes.func.isRequired,
    profile: PropTypes.object.isRequired
};

const mapStateToProps = state => ({
    profile: state.profile
});

export default connect(
    mapStateToProps,
    { updateProfile, getCurrentProfile }
)(withRouter(EditProfile));