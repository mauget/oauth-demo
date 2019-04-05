import React, {Component} from 'react';
import AuthContext from './AuthContext';

class AuthProvider extends Component {
    constructor(props) {
        super(props);

        this.state.isAuthenticated = false;
    }

    render() {
        return (
            <AuthContext.Provider
                value={{
                    signIn: () => {
                        this.setState({isAuthenticated: true} );
                    },
                    signOut: () => {
                        this.setState({isAuthenticated: false} );
                    }
                }}
            >
                {this.props.children}
            </AuthContext.Provider>
        );
    }
}