import React, {Component} from 'react';
import PropTypes from 'prop-types';
import {Button, Modal} from "react-bootstrap";

class LoginPopup extends Component {
    constructor(props) {
        super(props);

        this.state = {
            show: false,
            userId: '',
            password: ''
        };

        this.handleShow = this.handleShow.bind(this);
        this.handleClose = this.handleClose.bind(this);
        this.handleSignIn = this.handleSignIn.bind(this);
        this.onChangeUser = this.onChangeUser.bind(this);
        this.onChangePassword = this.onChangePassword.bind(this);
    }

    handleClose() {
        this.setState({show: false});
    }

    handleShow() {
        this.setState({show: true});
    }

    handleSignIn(ev) {
        ev.preventDefault();

        const {userId, password} = this.state;
        this.props.signIn({userId, password});

        this.handleClose();
    }

    onChangeUser(ev){
        this.setState({userId: ev.currentTarget.value});
    }

    onChangePassword(ev){
        this.setState({password: ev.currentTarget.value});
    }

    render() {
        return (
            <>
                <Button variant="primary" onClick={this.handleShow}>
                    Login
                </Button>

                <Modal show={this.state.show} onHide={this.handleClose}>
                    <form onSubmit={this.handleSignIn}>
                        <Modal.Header closeButton>
                            <Modal.Title>Log in with Google</Modal.Title>
                        </Modal.Header>

                        <Modal.Body>
                            <div className="form-group">
                                <label htmlFor="idUser">Google User</label>
                                <input type="text"  onChange={this.onChangeUser}
                                       className="form-control" id="idUser"/>
                            </div>
                            <div className="form-group">
                                <label htmlFor="isPassword">Password</label>
                                <input type="password"  onChange={this.onChangePassword}
                                       className="form-control" id="isPassword"/>
                            </div>
                        </Modal.Body>

                        <Modal.Footer>
                            <Button variant="secondary" onClick={this.handleClose}>Cancel</Button>
                            <Button type={"submit"}  variant="primary" >Login</Button>
                        </Modal.Footer>
                    </form>
                </Modal>
            </>
        )
    }
}

LoginPopup.propTypes = {
    signIn: PropTypes.func.isRequired,
};

export default LoginPopup;
