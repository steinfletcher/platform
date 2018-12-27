import {FETCH_USER, FETCH_USER_SUCCESS, FetchUserActions} from "./home.types";
import {FetchState} from "../api/http";

export interface HomeState {
    email?: string
    fetchState: FetchState
}

export const initialHomeState: HomeState = {
    email: undefined,
    fetchState: FetchState.UNKNOWN
};

export const homeReducer = (state: HomeState, action: FetchUserActions): HomeState => {
    switch (action.type) {
        case FETCH_USER:
            return {
                ...state,
                fetchState: FetchState.PENDING
            };
        case FETCH_USER_SUCCESS:
            return {
                ...state,
                fetchState: FetchState.SUCCESS,
                email: action.email
            };
        default:
            return state;
    }
};
