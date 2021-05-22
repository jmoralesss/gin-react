import axios, { AxiosResponse } from 'axios';
import { Endpoint} from './Endpoint.enum'

let BaseInstance = axios.create({
	baseURL: Endpoint.base,
	timeout: 15000,
});

const responseSuccessInteceptor = function(response: any): AxiosResponse<any> {
    const token = response.data.token;
    
    if(token !== undefined && token !== null) {
        sessionStorage.setItem('token', token);
    }
    
    return response;
}

const requestInteceptor = function(request: any): AxiosResponse<any> {
    let common = request.headers.common;
    common["Authorization"] =  sessionStorage.getItem("token");
    
    return request;
}

BaseInstance.interceptors.response.use(responseSuccessInteceptor);
BaseInstance.interceptors.request.use(requestInteceptor);

export {BaseInstance}

export type Response<T> = {
    data: T
}