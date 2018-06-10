// Framework
import React, {Component} from 'react';

// Style
import style from "./Summary.scss";

class Summary extends Component {
    constructor(props){
        super(props);
        this.summary = props.data;
    }

    render(){
        return (
          <div className={style.summary}>
              <h2>Summary</h2>
              <p>Guessed Days Left: {this.summary.days}</p>
              <p>Guessed Weeks Left: {this.summary.weeks}</p>
              <p>Guessed Date to Finish: {this.summary.last_day}</p>
              <p>Guessed Week to Finish: {this.summary.last_week}</p>
          </div>
        );
    }
}

export default Summary;