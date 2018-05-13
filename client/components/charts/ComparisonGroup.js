// Framework
import React, {Component} from 'react';
import { connect } from "react-redux";
import axios from "axios/index";

const mapStateToProps = state => {
    return {
        dates: state.comparisonSelector
    };
};

class ComparisonGroupComp extends Component{
    constructor(props){
        super(props);
        this.countPropsDates = 0;
        this.dataRecieved = false;
        this.state = {
            projectId: props.projectId,
            data: null
        }
    }

    shouldComponentUpdate(nextProps, nextState){
        let lastDates = nextProps.dates[nextProps.dates.length - 1];
        if(lastDates.left === "" || lastDates.right === ""){
            return false;
        }

        if(nextProps.dates.length === this.countPropsDates){
            if(this.dataRecieved === false && nextState.data !== null) {
                this.dataRecieved = true;
                return true;
            }
            return false;
        } else {
            this.dataRecieved = false;
            this.countPropsDates = nextProps.dates.length;
        }

        return true;
    }

    componentDidUpdate(){
        if(this.dataRecieved === true){
            return;
        }
        let dates = this.getComparisonSelectorDates();
        axios.get(`/data/stats/comparison/${this.state.projectId}/${dates.left}/${dates.right}`).then(res => {
            let newState = Object.assign({}, this.state);
            newState.data = res.data;
            this.setState(newState);
        });
    }

    getComparisonSelectorDates(){
        return this.props.dates[this.props.dates.length - 1];
    }

    render(){
        if(this.state.data === null){
            return(
                <p>Select the dates for the two versions of project {this.state.projectId}.</p>
            );
        }

        return(
          <p>Data received</p>
        );
    }
}

const ComparisonGroup = connect(mapStateToProps)(ComparisonGroupComp);
export default ComparisonGroup;