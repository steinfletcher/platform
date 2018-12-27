import * as React from 'react';

interface HomePageProps {
    email?: string
}

export const HomePage: React.FunctionComponent<HomePageProps> = props => (
    <div className="text-center">
        <h2>Home Page</h2>
        {
            props.email
                ? <div>Hi {props.email}</div>
                : null
        }
    </div>
);
