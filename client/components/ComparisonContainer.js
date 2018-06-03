// Framework
import React, {Component} from 'react';
import { connect } from "react-redux";
import axios from "axios/index";

// Style
import style from "./ComparisonContainer.scss";

// Components
import ComparisonGroup from "./charts/ComparisonGroup";
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
        this.chartType = props.type;
        this.projectId = props.project;
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
              <div className={style.dateSelector}>
                  <label>Dates to compare: </label>
                  {children}
              </div>
              <ComparisonGroup key={this.id + "-group"} id={this.id + "-group"} projectId={this.projectId}/>
          </div>
        );
    }
}

const ComparisonContainer = connect(mapStateToProps)(ComparisonContainerComp);
export default ComparisonContainer;