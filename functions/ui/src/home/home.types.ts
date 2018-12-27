
export const FETCH_USER = 'FETCH_USER';
export const FETCH_USER_SUCCESS = 'FETCH_USER_SUCCESS';

export class FetchUser {
    readonly type = FETCH_USER;
}

export class FetchUserSuccess {
    readonly type = FETCH_USER_SUCCESS;
    constructor(public email: string) {}
}

export type FetchUserActions = FetchUser | FetchUserSuccess
