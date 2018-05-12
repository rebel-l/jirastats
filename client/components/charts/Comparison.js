// Framework
import React, {Component} from 'react';
import { connect } from "react-redux";
import axios from "axios/index";

// Components
import DateSelector from "./../DateSelector";

const mapStateToProps = state => {
    return {
        data: state.chartContainer
    };
};

class ComparisonComp extends Component {
    constructor(props){
        super(props);
        this.id = "comparison-container";
        this.chartType = props.chartType;
        this.projectId = props.projectId;
        this.state = {
            dates: []
        };
    }

    componentDidMount(){
        axios.get(`/data/tickets/dates/${this.projectId}`).then(res => {
            this.setState({
                dates: res.data
            });
        });
    }

    getDateSelector(version, data){
        let selectorId = "date-selector";
        return (
            <DateSelector key={selectorId + "-" + version} id={selectorId + "-" + version} version={version} data={data}/>
        );
    }

    render(){
        let children = [];
        if(this.state.dates.length > 0){
            children.push((
                this.getDateSelector("left", this.state.dates)
            ));
            children.push((
                this.getDateSelector("right", this.state.dates)
            ));
        }

        return (
          <div key={this.id} id={this.id}>
              <div>
                  <label>Dates to compare: </label>
                  {children}
              </div>
              <div>Chart</div>
          </div>
        );
    }
}

const Comparison = connect(mapStateToProps)(ComparisonComp);
export default Comparison;