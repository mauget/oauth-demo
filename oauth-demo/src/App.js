import React, {Component, Fragment} from 'react';
import './App.css';
import Login from './login/Login'
import {BrowserRouter, Route, Switch} from "react-router-dom";
import PrivateRoute from './components/PrivateRoute'


class App extends Component {

    // TODO tie this to token presence in React context
    isAuthenticated = () => false;

    renderNeedLogin = () => <div style={{color: "red", backgroundColor: "white", padding: "1rem"}}>
        Please, please, please ... log in!</div>;

    renderAuthenticated = () => <div style={{color: "green", backgroundColor: "white", padding: "1rem"}}>
        You&apos;re a well-behaved, nice authentic user!</div>;

    showAuthState = () => this.isAuthenticated() ? this.renderAuthenticated() : this.renderNeedLogin();

    render() {

        const Main = () =>
            <Switch>
                <Route exact path="/" component={this.showAuthState}/>}/>
                <PrivateRoute path="/login" component={Login}/>

                <PrivateRoute component={this.renderNeedLogin}/>
            </Switch>;


        const Layout = () =>
            <Fragment>
                <BrowserRouter>
                    <Main/>
                </BrowserRouter>
            </Fragment>;

        return (
            <div className="App">
                <header className="App-header">
                    <Layout/>
                </header>
            </div>
        );
    }
}

export default App;
