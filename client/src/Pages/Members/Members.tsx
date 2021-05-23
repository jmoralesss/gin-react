import React from 'react';
import { MemberService } from '../../Services/Member.service'
import './Members.css'
import { toast } from 'react-toastify';
import { MembersState } from '../../Types/MembersState.type'
import { Member } from '../../Types/Member.type'

export class Members extends React.Component<{}, MembersState> {
    constructor(props: any) {
        super(props)

        this.state = {
            members: []
        }

        toast.configure()
    }

    componentDidMount(): void {
        MemberService.getMembers().then((res: any) => {
            const members: Member[] = res.data.map((member: any) => ({ email: member.Email, insertDate: new Date(member.InsertDate).toLocaleString() }))

            this.setState({ members: members })
        })
    }

    logout(): void {
        sessionStorage.removeItem('token')

        window.location.assign("/");
    }

    renderTableData() {
        return this.state.members.map((member, index) => {
            return (
                <tr key={index}>
                    <td>{member.email}</td>
                    <td>{member.insertDate}</td>
                </tr>
            )
        })
    }

    render(): JSX.Element {
        return (
            <div className="home-container">
                <div className="logout-container">
                    <button onClick={this.logout}>Logout</button>
                </div>
                <table>
                    <thead>
                        <tr>
                            <th>Email</th>
                            <th>Registration Date</th>
                        </tr>
                    </thead>
                    <tbody>
                        {this.renderTableData()}
                    </tbody>
                </table>
            </div>
        )
    }
}