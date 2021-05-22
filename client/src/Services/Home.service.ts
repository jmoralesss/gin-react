import { AxiosResponse } from 'axios';
import { FormEntryState } from '../Types/FormEntry.type'
import { Endpoint} from './Endpoint.enum'
import { BaseInstance } from './Base.service'
import { RegistrationResponse } from '../Types/RegistrationResponse.type';

const responseBody = (response: AxiosResponse) => response.data;

const requests = {
	post: (url: string, body: {}) => BaseInstance.post(url, body).then(responseBody)
};

export const HomeService = {
	login: (post: FormEntryState): Promise<boolean> => requests.post(Endpoint.login, post),
    logout: (post = {}): Promise<boolean> => requests.post(Endpoint.logout, post),
    register: (post: FormEntryState): Promise<RegistrationResponse> => requests.post(Endpoint.register, post),
};