import React from 'react';
import { MemberService } from '../../Services/Member.service'
import './Home.css'
import { toast } from 'react-toastify';
import { HomeState } from '../../Types/Home.type'

export class Home extends React.Component<{}, HomeState> {
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