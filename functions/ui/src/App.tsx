import React from 'react';

import Container from "react-bootstrap/lib/Container";
import {Header} from "./components/header/Header";
import {BrowserRouter as Router, Route} from "react-router-dom";
import {HomeContainer} from "./home/Home.container";
import {LoginContainer} from "./login/Login.container";
import {StoreProvider} from "./store";

export const App = () => (
    <Container>
        <StoreProvider>
            <Router>
                <>
                    <Header/>
                    <Route exact path="/" component={HomeContainer}/>
                    <Route path="/login" component={LoginContainer}/>
                </>
            </Router>
        </StoreProvider>
    </Container>
);
