import React, {Component} from 'react';
import './App.css';
import AuthExample from './AuthExample';

class App extends Component {
    render() {
        return (
            <div className="App">
                <header className="App-header">
                    <AuthExample/>
                </header>
            </div>
        );
    }
}

export default App;
