// Framework
import React, {Component} from "react";

class Footer extends Component {
    constructor(props){
        super(props);
        this.licenceLink = "https://www.gnu.org/licenses/gpl-3.0";
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
            <div key={"footer"} id={"footer"}>
                <p>
                    Copyright © {Footer.getYear()} by <a href={this.profileLink}>Lars Gaubisch</a>
                </p>
                <p>
                    Published under <a href={this.licenceLink}>GNU GPL.3.0</a>
                </p>
            </div>
        );
    }
}



export default Footer;