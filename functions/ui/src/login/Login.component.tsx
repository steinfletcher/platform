import React, {FormEvent} from "react";
import Button from 'react-bootstrap/lib/Button';
import Form from 'react-bootstrap/lib/Form';

type LoginFormProps = {
    username: string
    password: string
    disabled: boolean
    loginPending: boolean
    errorMessage?: string
    updateUsername: (username: string) => void
    updatePassword: (password: string) => void
    onCreateSession: () => void
}

export const LoginForm: React.FunctionComponent<LoginFormProps> = props => {
    return (
        <Form onSubmit={(e: React.FormEvent<HTMLFormElement>) => {
            e.preventDefault();
            props.onCreateSession();
        }}>
            <Form.Group controlId="formBasicUsername">
                <Form.Label>Username</Form.Label>
                <Form.Control onChange={(e: FormEvent<HTMLInputElement>) => props.updateUsername(e.currentTarget.value)}
                              type="text"
                              placeholder="Enter username"
                              defaultValue={props.username}/>
            </Form.Group>

            <Form.Group controlId="formBasicPassword">
                <Form.Label>Password</Form.Label>
                <Form.Control onChange={(e: FormEvent<HTMLInputElement>) => props.updatePassword(e.currentTarget.value)}
                              type="password"
                              placeholder="Password"
                              defaultValue={props.password}/>

                {props.errorMessage ? <Form.Text data-test-id="incorrectCredentials" className="text-danger">{props.errorMessage}</Form.Text> : null }
            </Form.Group>

            <Button data-test-id="submitBtn" variant="primary" type="submit" disabled={props.disabled}>
                {props.loginPending ? 'Loading...' : 'Submit'}
            </Button>
        </Form>
    );
};
