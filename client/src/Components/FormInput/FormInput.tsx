import React from 'react';
import './FormInput.css';
import { IFormEntry } from '../../Interfaces/IFormEntry'

export class FormEntry extends React.Component<IFormEntry> {
    render() {
        return (
            <span>
                <input onChange={this.props.handleEmailChange} placeholder="Email" type="text"></input>
                <input onChange={this.props.handlePasswordChange}  placeholder="Password" type="password"></input>
            </span>
        )
    }
}