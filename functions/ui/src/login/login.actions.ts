import {History} from 'history'
import {Dispatch} from "react";
import * as api from "../api";
import {LoginState} from "./login.state";
import {url} from "../urls";

export enum ActionType {
    UpdateUsername = "UPDATE_USERNAME",
    UpdatePassword = "UPDATE_PASSWORD",
    LoginSuccess = "LOGIN_SUCCESS",
    IncorrectCredentials = "INCORRECT_CREDENTIALS",
    UnknownError = "UNKNOWN_ERROR",
    SubmitCredentials = "SUBMIT_CREDENTIALS"
}

export type UpdateFieldAction = {
    type: ActionType.UpdatePassword | ActionType.UpdateUsername
    value: string
}

export type LoginSuccessAction = {
    type: ActionType.LoginSuccess
}

export type IncorrectCredentialsAction = {
    type: ActionType.IncorrectCredentials,
    errorMessage: string
}

export type UnknownErrorAction = {
    type: ActionType.UnknownError
}

export type SubmitCredentialsAction = {
    type: ActionType.SubmitCredentials
}

export type LoginActionTypes =
    UpdateFieldAction |
    LoginSuccessAction |
    IncorrectCredentialsAction |
    UnknownErrorAction |
    SubmitCredentialsAction

export const updateUsernameAction = (dispatch: Dispatch<UpdateFieldAction>) =>
    (username: string) => dispatch({type: ActionType.UpdateUsername, value: username});

export const updatePasswordAction = (dispatch: Dispatch<UpdateFieldAction>) =>
    (username: string) => dispatch({type: ActionType.UpdatePassword, value: username});

export const createSessionAction = (dispatch: Dispatch<LoginSuccessAction | IncorrectCredentialsAction |
    UnknownErrorAction | SubmitCredentialsAction>, state: LoginState, history: History) =>
    async (): Promise<void> => {
        dispatch({type: ActionType.SubmitCredentials});

        const result = await api.createSession({
            username: state.username,
            password: state.password
        });

        if (result.isSuccess()) {
            dispatch({type: ActionType.LoginSuccess});
            history.push(url.home);
        } else if (result.isClientError()) {
            dispatch({type: ActionType.IncorrectCredentials, errorMessage: result.error.description});
        } else {
            dispatch({type: ActionType.UnknownError});
        }
    };
