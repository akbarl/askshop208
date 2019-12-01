import React, { Component } from 'react';
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link
} from "react-router-dom";

import Product from '../../product'
import HomePage from '../../homepage'
class Header extends Component {
    render() {
        return (
            <Router>
                <div>
                    <ul>
                        <li>
                            <Link to="/">Homepage</Link>
                        </li>
                        <li>
                            <Link to="/product">Product</Link>
                        </li>
                    </ul>

                    <hr />
                    <Switch>
                        <Route exact path="/">
                            <HomePage />
                        </Route>
                        <Route exact path="/product">
                            <Product />
                        </Route>
                    </Switch>
                </div>
            </Router>
        );
    }
}

export default Header;