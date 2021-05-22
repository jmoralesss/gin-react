import { ChangeEventHandler } from 'react';

export interface IFormEntry {
    handleEmailChange: ChangeEventHandler<HTMLInputElement>
    handlePasswordChange: ChangeEventHandler<HTMLInputElement>
}
