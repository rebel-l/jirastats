// Framework
import axios from "axios/index";
import React, {Component} from 'react';
import { connect } from "react-redux";

// Actions
import ChartContainerAction from "./../actions/ChartContainer";

// Components
import ChartLineColumn from "./charts/LineColumn";
import PieChartTable from "./charts/PieChartTable";
import Forecast from "./charts/Forecast";

// Constants
import {CHARTTYPE_FORECAST} from "./../constants/ChartTypes";
import {CHARTTYPE_OPENTICKETS} from "./../constants/ChartTypes";
import {CHARTTYPE_PROGRESS} from "./../constants/ChartTypes";
import {CHARTTYPE_SPEED} from "./../constants/ChartTypes";

const mapStateToProps = state => {
    return {
        chartButton: state.chartButton,
        project: state.project
    };
};

const mapDispatchToProps = dispatch => {
    return {
        chartContainer: chartContainer => dispatch(ChartContainerAction(chartContainer))
    };
};

class ChartContainerComp extends  Component {
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
                    id: "LineChart-" + key,
                    type: "line",
                    title: "Progress",
                    xAxis: "Date",
                    yAxis: "Number of Tickets"
                };
                child = (
                    <ChartLineColumn key={key} options={options}/>
                );
                break;
            case CHARTTYPE_SPEED:
                options = {
                    id: "ColumnChart-" + key,
                    type: "column",
                    title: "Speed",
                    xAxis: "Week",
                    yAxis: "Number of Tickets"
                };
                child = (<ChartLineColumn key={key} options={options} />);
                break;
            case CHARTTYPE_OPENTICKETS:
                child = (<PieChartTable key={key} type={chartType} project={project}/>);
                break;
            case CHARTTYPE_FORECAST:
                child = (
                    <Forecast key={key} type={chartType} project={project}/>
                );
                break;
            default:
                child = (<p>"{chartType}" not implemented yet.</p>);
                break
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