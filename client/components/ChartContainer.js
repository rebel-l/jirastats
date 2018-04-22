// Framework
import axios from "axios/index";
import React, {Component} from 'react';
import { connect } from "react-redux";

// Actions
import ChartContainerAction from "./../actions/ChartContainer";

// Components
import ChartLine from "./charts/Line";
import ChartColumn from "./charts/Column";

// Constants
import {CHARTTYPE_PROGRESS} from "./../constants/ChartTypes";
import {CHARTTYPE_SPEED} from "./../constants/ChartTypes";
import {CHARTTYPE_OPENTICKETS} from "./../constants/ChartTypes";

const mapStateToProps = state => {
    return {
        chartButton: state.chartButton,
        // chartContainer: state.chartContainer,
        project: state.project
    };
};

const mapDispatchToProps = dispatch => {
    return {
        chartContainer: chartContainer => dispatch(ChartContainerAction(chartContainer))
    };
};

class ChartContainerComp extends  Component {
    constructor(props){
        super(props);
    }

    componentDidUpdate () {
        let chartType = this.getChartType();
        let project = this.getProjectId();
        axios.get(`/data/stats/${chartType}/${project}`).then(res => {
            this.props.chartContainer(res.data);
        });
    }

    getChartType(){
        return this.props.chartButton[this.props.chartButton.length - 1];
    }

    getProjectId(){
        return this.props.project[this.props.project.length - 1];
    }

    render(){
        let id = 'ChartContainer';
        let chartType = this.getChartType();
        let project = this.getProjectId();
        let key = project + '-' + chartType;
        let child, options;
        switch (chartType) {
            case CHARTTYPE_PROGRESS:
                options = {
                    title: 'Progress',
                    xAxis: 'Date',
                    yAxis: 'Number of Tickets'
                };
                child = (<ChartLine key={key} type={chartType} project={project} options={options} />);
                break;
            case CHARTTYPE_SPEED:
                options = {
                    title: 'Speed',
                    xAxis: 'Week',
                    yAxis: 'Number of Tickets'
                };
                child = (<ChartColumn key={key} type={chartType} project={project} options={options} />);
                break;
        }

        return(
            <div key={id} id={id}>
                {child}
            </div>
        );
    }
}

const ChartContainer = connect(mapStateToProps, mapDispatchToProps)(ChartContainerComp);
export default ChartContainer