// Framework
import React, {Component} from 'react';
import { connect } from "react-redux";
import axios from "axios/index";

// Components
import DateSelector from "./DateSelector";

const mapStateToProps = state => {
    return {
        data: state.chartContainer
    };
};

class ComparisonContainerComp extends Component {
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

    static getDateSelector(version, data){
        let selectorId = "date-selector";
        return (
            <DateSelector key={selectorId + "-" + version} id={selectorId + "-" + version} version={version} data={data}/>
        );
    }

    render(){
        let children = [];
        if(this.state.dates.length > 0){
            children.push((
                ComparisonContainerComp.getDateSelector("left", this.state.dates)
            ));
            children.push((
                ComparisonContainerComp.getDateSelector("right", this.state.dates)
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

const ComparisonContainer = connect(mapStateToProps)(ComparisonContainerComp);
export default ComparisonContainer;