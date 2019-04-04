import React, {Component} from 'react';
import {Button, Modal} from "react-bootstrap";

class LoginPopup extends Component {
    constructor(props) {
        super(props);

        this.state = {
            show: false,
        };

        this.handleShow = this.handleShow.bind(this);
        this.handleClose = this.handleClose.bind(this);
        this.handleAuthentication = this.handleAuthentication.bind(this);
    }

    handleClose() {
        this.setState({show: false});
    }

    handleShow() {
        this.setState({show: true});
    }

    handleAuthentication() {
        this.props.authenticate();
        this.handleClose();
    }

    render() {
        return (
            <>

                <Button variant="primary" onClick={this.handleShow}>
                    Login
                </Button>

                <Modal show={this.state.show} onHide={this.handleClose}>
                    <Modal.Header closeButton>
                        <Modal.Title>Modal heading</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>Woohoo, you're reading this text in a modal!</Modal.Body>
                    <Modal.Footer>
                        <Button variant="secondary" onClick={this.handleClose}>
                            Cancel
                        </Button>
                        <Button variant="primary" onClick={this.handleAuthentication}>
                            Login
                        </Button>
                    </Modal.Footer>
                </Modal>
            </>
        )
    }
}

export default LoginPopup;
