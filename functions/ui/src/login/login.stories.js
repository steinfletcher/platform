import React from 'react';
import {storiesOf} from '@storybook/react';
import Container from 'react-bootstrap/lib/Container';
import 'bootstrap/dist/css/bootstrap.min.css'

import {noop} from '../shared/lang'
import {LoginForm} from './Login.component.tsx'
import {incorrectCredentialsMessage} from './login.state.ts'

const actionProps = {
    updateUsername: noop,
    updatePassword: noop,
    onCreateSession: noop,
};

storiesOf('Login Form', module)
    .add('default', () =>
        <Container fluid>
            <LoginForm
                disabled={false}
                username={"a@b.com"}
                password={"Password123"}
                loginPending={false}
                {...actionProps}
            />
        </Container>
    )
    .add('disabled', () =>
        <Container fluid>
            <LoginForm
                disabled={true}
                email={""}
                password={""}
                loginPending={false}
                {...actionProps}
            />
        </Container>
    )
    .add('login pending', () =>
        <Container fluid>
            <LoginForm
                disabled={true}
                username={"a@b.com"}
                password={"Password123"}
                loginPending={true}
                {...actionProps}
            />
        </Container>
    )
    .add('incorrect credentials', () =>
        <Container fluid>
            <LoginForm
                username={"a@b.com"}
                password={"Password123"}
                errorMessage={incorrectCredentialsMessage}
                {...actionProps}
            />
        </Container>
    );
