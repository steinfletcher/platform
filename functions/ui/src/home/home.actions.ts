import {Dispatch} from "react";
import {History} from "history";
import * as api from "../api";
import {url} from "../urls";
import {FetchUser, FetchUserActions, FetchUserSuccess} from "./home.types";

export const fetchProfileAction = (dispatch: Dispatch<FetchUserActions>, history: History) =>
    async (): Promise<void> => {
        dispatch(new FetchUser());

        const result = await api.fetchProfile();
        if (result.isSuccess()) {
            dispatch(new FetchUserSuccess(result.data.username));
        } else if (result.isUnauthorized()) {
            history.push(url.login);
        }
    };
