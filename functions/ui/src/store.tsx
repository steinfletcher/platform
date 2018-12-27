import * as React from 'react';
import {Dispatch, Reducer, useContext, useReducer} from 'react';
import {homeReducer, HomeState, initialHomeState} from "./home/home.state";
import {loginReducer, LoginState, initialLoginState} from "./login/login.state";

export interface State {
    home: HomeState
    login: LoginState
}

const initialState: State = {
    home: initialHomeState,
    login: initialLoginState
};

interface Context {
    state: State,
    dispatch: Dispatch<any>
}

type Reducers = {
    [key: string]: Reducer<any, any>
}

const Store = React.createContext<Context | null>(null);

const reducers = {
    home: homeReducer,
    login: loginReducer
};

export const StoreProvider = ({children}: { children: JSX.Element }) => {
    const store = createStore(reducers);
    return <Store.Provider value={store}>{children}</Store.Provider>;
};

function createStore(reducers: Reducers): Context {
    const [state, dispatch] = useReducer(combineReducers(reducers), initialState);
    return {state, dispatch};
}

export function useStore(): Context {
    return useContext(Store)!;
}

function combineReducers(reducers: Reducers): Reducer<any, any> {
    return (state: any, action: Dispatch<any>): any => {
        return Object.keys(reducers).reduce(
            (nextState: any, key: string): any => {
                nextState[key] = reducers[key](
                    state[key],
                    action
                );
                return nextState;
            },
            initialState
        );
    };
}
