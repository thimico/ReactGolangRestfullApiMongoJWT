import { combineReducers } from 'redux';
import authReducer from './authReducer';
import errReducer from './errorReducer';
import profileReducer from './profileReducer';

export default combineReducers({
    auth: authReducer,
    errors: errReducer,
    profile: profileReducer
})

