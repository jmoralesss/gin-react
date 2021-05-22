import React from 'react'
import { FormEntry } from '../../Components/FormInput/FormInput'
import { FormEntryProps, FormEntryState } from '../../Types/FormEntry.type'
import { HomeService } from '../../Services/Home.service'
import { toast } from 'react-toastify';
import "./Register.css"

export class Register extends React.Component<FormEntryProps, FormEntryState> {
    constructor(props: any) {
        super(props)

        toast.configure()

        this.state = {
            email: '',
            password: ''
        }
    }

    handleEmailChange = (event: any) => this.setState({ email: event.target.value })
    handlePasswordChange = (event: any) => this.setState({ password: event.target.value })

    handleRegistration(event: any) {
        event.preventDefault();

        const passedValidation = this.validate(this.state)
        if(!passedValidation) {
            return;
        }

        HomeService.register(this.state).then(res => {
            toast.success("Account successfully created!");

            this.props.history.push("/");
        }, err => {
            toast.error(err.response.data.error);
        })
    }

    validate(state: FormEntryState): boolean {
        const emailRegex = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
        if(!emailRegex.test(String(state.email).toLowerCase())) {
            toast.error("Please provide a valid email")
            return false;
        }
        
        if(state.password === "") {
            toast.error("Please provide a password")
            return false;
        }

        return true;
    }

    render(): JSX.Element {
        return (
            <div id="registration-container">
                <form onSubmit={this.handleRegistration.bind(this)}>
                    <FormEntry handleEmailChange={this.handleEmailChange} handlePasswordChange={this.handlePasswordChange}></FormEntry>
                    <div id="register-actions">
                        <input type="submit" value="Register"></input>
                    </div>
                </form>
            </div>
        )
    }
}