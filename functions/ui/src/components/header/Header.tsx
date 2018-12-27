import React from 'react';

import Navbar from 'react-bootstrap/lib/Navbar';
import Nav from 'react-bootstrap/lib/Nav';
import {Link} from "react-router-dom";

export const Header: React.FunctionComponent = () => (
    <Navbar sticky="top" variant="dark">
        <Navbar.Brand><Link to="/">Sample app</Link></Navbar.Brand>
        <Nav className="mr-auto">
            <Nav.Link><Link to="/login">Login</Link></Nav.Link>
        </Nav>
    </Navbar>
);
