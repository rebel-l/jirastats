// Framework
import React, {Component} from 'react';
import { connect } from "react-redux";

const mapStateToProps = state => {
    return {
        dates: state.comparisonSelector
    };
};

class ComparisonGroupComp extends Component{
    constructor(props){
        super(props);
        this.state = {
            projectId: props.projectId
        }
    }

    componentDidUpdate(){
        let dates = this.getComparisonSelectorDates();
        if(dates.left === "" || dates.right === ""){
            return;
        }

        // TODO: do ajax call (but endpoint is needed first)
    }

    getComparisonSelectorDates(){
        return this.props.dates[this.props.dates.length - 1];
    }

    render(){
        return(
            <p>Select the dates for the two versions of project {this.state.projectId}.</p>
        );
    }
}

const ComparisonGroup = connect(mapStateToProps)(ComparisonGroupComp);
export default ComparisonGroup;