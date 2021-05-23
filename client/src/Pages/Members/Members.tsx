import React from 'react';
import { MemberService } from '../../Services/Member.service'
import './Members.css'
import { toast } from 'react-toastify';
import { MembersState } from '../../Types/MembersState.type'

export class Members extends React.Component<{}, MembersState> {
    constructor(props: any) {
        super(props)

        this.state = {
            members: []
        }

        toast.configure()
    }

    componentDidMount(): void {
        MemberService.getMembers().then(res => {
            this.setState({ members: res.data})
        })
        console.log(this.state.members)
    }

    logout() {
        sessionStorage.removeItem('token')

        window.location.assign("/");
    }

    render(): JSX.Element {
        return (
            <div className="home-container">
                <div className="logout-container">
                    <button onClick={this.logout}>Logout</button>
                </div>
                <h1>Successfully logged in</h1>
            </div>
        )
    }
}