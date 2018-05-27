// Framework
import React, {Component} from "react";

// Style
import style from "./Footer.scss";

class Footer extends Component {
    constructor(props){
        super(props);
        this.licenseLink = "https://www.gnu.org/licenses/gpl-3.0";
        this.profileLink = "https://www.linkedin.com/in/lars-gaubisch-a6b1a423/";
    }

    static getYear(){
        let firstYear = 2018;
        let actualYear = new Date().getFullYear();

        if(firstYear === actualYear) {
            return firstYear;
        }
        return firstYear + " - " + actualYear;
    }

    render(){
        return(
            <div key={"footer"} id={"footer"} className={style.footer}>
                <p>
                    Copyright © {Footer.getYear()} by <a href={this.profileLink}>Lars Gaubisch</a>
                </p>
                <p>
                    Licensed under <a href={this.licenseLink}>GNU GPL.3.0</a>
                </p>
            </div>
        );
    }
}

export default Footer;