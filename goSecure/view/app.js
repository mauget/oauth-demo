(function () {

    // noinspection JSUnresolvedVariable,JSUnresolvedVariable
    class Demo extends React.Component {
        constructor(props) {
            super(props);
            this.state = {
                liked: false,
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


        logIn = async ()=> {
            const body = {userID: this.state.userID, password: this.state.password};


            try {
                const {data} =await api.post('login', body);
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


        render() {
            if (this.state.liked) {
                return 'You liked this.';
            }

            // noinspection JSXNamespaceValidation,JSUnresolvedFunction
            return (
                <div>
                    <button onClick={() => this.setState({liked: true})}>Like</button>
                    <button onClick={() => this.logMessage()}>Log Message</button>

                    <br/>
                    <br/>
                    <form onSubmit={this.logIn}>
                        <label htmlFor={"userID"}>User ID:
                            <input id={"userID"} onChange={this.onUserID} type={"text"}/>
                        </label>
                        <br/>
                        <label htmlFor={"password"}>Password:
                            <input id={"password"} type={"password"} onChange={this.onPassword} />
                        </label>
                        <br/>
                        <button type={"submit"}>Log In</button>
                    </form>
                    <br/>
                    <br/>

                    <button onClick={() => this.logOut()}>Log Out</button>
                </div>
            );
        }
    }

    const domContainer = document.querySelector('#root');
    // noinspection JSUnresolvedVariable
    ReactDOM.render(React.createElement(Demo), domContainer);

    //===========================================================================
    // Client service:

    const URL = `https://${location.hostname}:443/api/`;

    const api = axios.create({
        baseURL: URL,
        headers: {'Content-Type': 'application/json'},
        withCredentials: true
    });

})();