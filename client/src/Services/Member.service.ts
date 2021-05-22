import { AxiosResponse } from 'axios';
import { Endpoint} from './Endpoint.enum'
import { BaseInstance } from './Base.service'
import { Member } from '../Types/Member.type';
import { Response } from '../Services/Base.service'

const responseBody = (response: AxiosResponse) => response.data;

const requests = {
	post: (url: string, body: {}) => BaseInstance.post(url, body).then(responseBody),
	get: (url: string) => BaseInstance.get(url).then(responseBody)
};

export const MemberService = {
    getMembers: (): Promise<Response<any[]>> => requests.get(Endpoint.memberGet),
};
