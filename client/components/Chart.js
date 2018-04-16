// Framework
import axios from "axios/index";
import React, {Component} from 'react';
import { connect } from "react-redux";
import Highcharts from 'highcharts';

// Constants
import {CHARTTYPE_PROGRESS} from "./../constants/ChartTypes";
import {CHARTTYPE_SPEED} from "./../constants/ChartTypes";

const mapStateToProps = state => {
    return {
        chartData: state.chart,
        chartButton: state.chartButton,
        project: state.project
    };
};

class ChartComp extends Component {
    componentDidUpdate () {
        let chartType = this.props.chartButton[this.props.chartButton.length - 1];
        let project = this.props.project[this.props.project.length - 1];

        axios.get(`/data/stats/${chartType}/${project}`).then(res => {
            switch (chartType) {
                case CHARTTYPE_PROGRESS:
                    Highcharts.chart('chart', this.getChartOptions("line", res.data));
                    break;
                case CHARTTYPE_SPEED:
                    Highcharts.chart('chart', this.getChartOptions("column", res.data));
                    break;
            }
        });
    }

    getChartOptions(type, chartData) {
        return {
            chart: {
                type: type
            },
            title: {
                text: chartData.project_name
            },
            subtitle: {
                text: 'Progress'
            },
            xAxis: {
                title: {
                    text: "Date"
                },
                categories: chartData.categories
            },
            yAxis: {
                title: {
                    text: 'Number of Tickets'
                }
            },
            legend: {
                layout: 'vertical',
                align: 'right',
                verticalAlign: 'middle'
            },
            series: chartData.series,
            credits: false
        }
    }

    render() {
        let chartType = this.props.chartButton[this.props.chartButton.length - 1];
        switch (chartType) {
            case CHARTTYPE_PROGRESS:
            case CHARTTYPE_SPEED:
                return (
                    <div key={"chartContainer"}>
                        <div id="chart" />
                    </div>
                );
            default:
                return (
                    <div key={"chartContainer"}>
                        <p>"{chartType}" not implemented yet.</p>
                    </div>
                )
        }

    }
}

const Chart = connect(mapStateToProps)(ChartComp);

export default Chart
