import React from 'react';

// this is the equivalent to the createStore method of Redux
// https://redux.js.org/api/createstore

const AuthContext = React.createContext({isAuthenticated: false});

export default AuthContext;
