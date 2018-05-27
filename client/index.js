// Frameworks
import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from "react-redux";
import store from "./store";

// Style
import "normalize-css";
import "./base.scss";

// Components
import App from './components/App';

ReactDOM.render(
    <Provider store={store}>
        <App/>
    </Provider>,
    document.getElementById('appInit')
);
