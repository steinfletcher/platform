import {ActionType, LoginActionTypes} from "./login.actions";
import {FetchState} from "../api/http";

export interface LoginState {
    username: string
    password: string
    disabled: boolean
    errorMessage?: string
    fetchState: FetchState
}

export const initialLoginState: LoginState = {
    username: "",
    password: "",
    disabled: true,
    fetchState: FetchState.UNKNOWN
};

export const incorrectCredentialsMessage = 'Incorrect credentials. Please try again'; // TODO move to backend

export const loginReducer = (state: LoginState, action: LoginActionTypes): LoginState => {
    switch (action.type) {
        case ActionType.UpdatePassword:
            return {
                ...state,
                password: action.value,
                disabled: !state.username || !action.value
            };
        case ActionType.UpdateUsername:
            return {
                ...state,
                username: action.value,
                disabled: !state.password || !action.value
            };
        case ActionType.LoginSuccess:
            return initialLoginState;
        case ActionType.IncorrectCredentials:
            return {
                ...state,
                errorMessage: incorrectCredentialsMessage,
                fetchState: FetchState.FAILURE
            };
        case ActionType.SubmitCredentials:
            return {
                ...state,
                errorMessage: undefined,
                fetchState: FetchState.PENDING
            };
        default:
            return state;
    }
};
