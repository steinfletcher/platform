import {apiUrl, Method, request, Result} from "./http";

const identityApi = (path: string) => apiUrl('identity', path);

interface CreateSessionRequest {
    username: string
    password: string
}

export const createSession = (requestModel: CreateSessionRequest): Promise<Result<void>> => {
    return request<void>(Method.POST, identityApi('/v1/session'), requestModel)
};
