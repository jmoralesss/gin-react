import React from 'react';
import { Link } from 'react-router-dom';
import { FormEntry } from '../../Components/FormInput/FormInput'
import { HomeService } from '../../Services/Home.service'
import { FormEntryState } from '../../Types/FormEntry.type'
import './Login.css'
import { toast } from 'react-toastify';

export class Login extends React.Component<{}, FormEntryState> {
    constructor(props: any) {
        super(props)

        toast.configure();

        this.state = {
            email: '',
            password: ''
        }
    }

    handleEmailChange = (event: any) => this.setState({ email: event.target.value })
    handlePasswordChange = (event: any) => this.setState({ password: event.target.value })

    handleLogin(event: any) {
        event.preventDefault();

        HomeService.login(this.state).then((res: any) => {
            window.location.assign("/members");
        }, err => {
            toast.error("Wrong email or password");
        })
    }

    render(): JSX.Element {
        return (
            <div id="login-container">
                <form onSubmit={this.handleLogin.bind(this)}>
                    <FormEntry handleEmailChange={this.handleEmailChange} handlePasswordChange={this.handlePasswordChange}></FormEntry>
                    <div className="actions">
                        <Link id="register-link" to='/register'>Register</Link>
                        <input type="submit" value="Login"></input>
                    </div>
                </form>
            </div>
        )
    }
}