(function () {


    // noinspection JSUnresolvedVariable,JSUnresolvedVariable
    class App extends React.Component {
        constructor(props) {
            super(props);

            // noinspection JSPotentiallyInvalidUsageOfThis
            this.state = {
                userID: null,
                password: null,
            };

        }

        logMessage = async () => {
            try {
                const {data} = await api.get(`testmsg`);
                console.log(`Server response: ${data}`);
            } catch (e) {
                console.error(e);
            }
        };


        logInHandler = async () => {
            const body = {userID: this.state.userID, password: this.state.password};


            try {
                const {data} = await api.post('login', body);
                console.log(`Login response: ${data}`);
            } catch (e) {
                console.error(e);
                alert('Login failed');
            }
        };

        logOut = async () => {
            await api.post('logout');
        };

        onUserID = (evt) => {
            this.setState({"userID": evt.target.value})
        };

        onPassword = (evt) => {
            this.setState({"password": evt.target.value})
        };

        onTestSession = async () => {
            try {
                const {data} = await api.get('session');
                console.log(`onTestSession response: ${data}`);
            } catch (e) {
                console.error(e);
                alert('onTestSession failed');
            }
        };


        render() {

            // noinspection JSXNamespaceValidation
            const renderLoginForm = (
                <form onSubmit={this.logInHandler}>
                    <label htmlFor={"userID"}>User ID:
                        <input id={"userID"} onChange={this.onUserID} type={"text"}/>
                    </label>
                    <br/>
                    <label htmlFor={"password"}>Password:
                        <input id={"password"} type={"password"} onChange={this.onPassword}/>
                    </label>
                    <br/>
                    <button type={"submit"}>Log In</button>
                </form>
            );

            const renderLogout = <button onClick={this.logOut}>Log Out</button>;

            const renderSessionTest = <button onClick={this.onTestSession}>Test Session</button>;

            const renderLogFromServer = <button onClick={() => this.logMessage()}>Log Message</button>;

            // noinspection JSXNamespaceValidation,JSUnresolvedFunction
            return (
                <div>
                    {renderLogFromServer}

                    <br/>
                    <br/>
                    {renderLoginForm}
                    <br/>
                    <br/>

                    {renderLogout}
                    <br/>
                    <br/>
                    {renderSessionTest}
                </div>
            );
        }
    }

    const domContainer = document.querySelector('#root');
    // noinspection JSUnresolvedVariable
    ReactDOM.render(React.createElement(App), domContainer);

    //===========================================================================
    // Client service:

    const URL = `https://${location.hostname}:${location.port}/api/`;

    const api = axios.create({
        baseURL: URL,
        headers: {'Content-Type': 'application/json'},
        withCredentials: true
    });

})();