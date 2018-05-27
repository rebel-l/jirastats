// Framework
import React, {Component} from "react";

// Style
import style from "./Header.scss";

class Header extends Component {
    render(){
        return(
            <header>
                <h1 className={style.title}>JIRA Project Statistics</h1>
            </header>
        );
    }
}

export default Header;