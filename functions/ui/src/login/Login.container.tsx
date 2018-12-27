import React from "react";
import {LoginForm} from "./Login.component";
import {createSessionAction, updatePasswordAction, updateUsernameAction} from "./login.actions";
import {FetchState} from "../api/http";
import {RouteComponentProps} from "react-router";
import {useStore} from "../store";

export const LoginContainer: React.FunctionComponent<RouteComponentProps<any>> = props => {
    const {state: {login}, dispatch} = useStore();

    console.log(login)

    return (
        <LoginForm
            disabled={login.disabled}
            username={login.username}
            password={login.password}
            errorMessage={login.errorMessage}
            loginPending={login.fetchState === FetchState.PENDING}
            updateUsername={updateUsernameAction(dispatch)}
            updatePassword={updatePasswordAction(dispatch)}
            onCreateSession={createSessionAction(dispatch, login, props.history)}
        />
    )
};
