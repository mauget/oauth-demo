import React, {Component} from "react";
import {BrowserRouter as Router, Link, Redirect, Route, withRouter} from "react-router-dom";
import LoginPopup from './LoginPopup';

const AuthExample = () => (
    <Router>
        <div>
            <AuthButton/>
            <ul>
                <li>
                    <Link to="/public">Public Page</Link>
                </li>
                <li>
                    <Link to="/protected">Protected Page</Link>
                </li>
            </ul>
            <Route path="/public" component={Public}/>
            <Route path="/login" component={Login}/>
            <PrivateRoute path="/protected" component={Protected}/>
        </div>
    </Router>
);


const authenticator = {
    isAuthenticated: false,
    authenticate(cb) {
        this.isAuthenticated = true;
        setTimeout(cb, 0);
    },
    signout(cb) {
        this.isAuthenticated = false;
        setTimeout(cb, 0);
    }
};


const AuthButton = withRouter(
    ({history}) =>
        authenticator.isAuthenticated ? (
            <p>
                Welcome!{" "}
                <button
                    onClick={() => {
                        authenticator.signout(() => history.push("/"));
                    }}
                >
                    Sign out
                </button>
            </p>
        ) : (
            <p>You are not logged in.</p>
        )
);


const PrivateRoute = ({component: Component, ...rest}) => (
    <Route
        {...rest}
        render={props =>
            authenticator.isAuthenticated ? (
                <Component {...props} />
            ) : (
                <Redirect
                    to={{
                        pathname: "/login",
                        state: {from: props.location}
                    }}
                />
            )
        }
    />
);


const Public = () => <h3>Public</h3>;

const Protected = () => <h3>Protected</h3>;

// ===========================================================
// Demo public/private modes:
// 1. Click the public page
// 2. Click the protected page
// 3. Log in
// 4. Click the back button, note the URL each time
// ===========================================================
class Login extends Component {
    constructor(props) {
        super(props);

        this.state = {redirectToReferrer: false};

        this.signIn = this.signIn.bind(this);
    }

    signIn({userId, password}) {

        console.log(`user ${userId}, password ${password}`);

        authenticator.authenticate(() => {
            this.setState({redirectToReferrer: true});
        });
    };

    render() {
        const {from} = this.props.location.state || {from: {pathname: "/public"}};
        const {redirectToReferrer} = this.state;

        if (redirectToReferrer) return <Redirect to={from}/>;

        return (
            <div>
                <p>You must log in to view the page at {from.pathname}</p>
                <LoginPopup signIn={this.signIn}/>
            </div>
        );
    }
}

export default AuthExample;
