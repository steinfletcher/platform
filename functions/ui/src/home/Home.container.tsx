import * as React from 'react';
import {useEffect} from 'react';
import {RouteComponentProps} from "react-router";
import {HomePage} from "./Home.component";
import {fetchProfileAction} from "./home.actions";
import {useStore} from "../store";

export const HomeContainer: React.FunctionComponent<RouteComponentProps<any>> = props => {
    const {state, dispatch} = useStore();

    useEffect(() => {
        fetchProfileAction(dispatch, props.history)()
    }, []);

    return <HomePage email={state.home.email}/>;
};
