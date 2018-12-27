import {Method, request, Result, apiUrl} from "./http";

const userApi = (path: string) => apiUrl('user', path);

interface FetchProfileResponse {
    username: string
}

export const fetchProfile = (): Promise<Result<FetchProfileResponse>> => {
    return request<FetchProfileResponse>(Method.GET, userApi('/v1/user'));
};
