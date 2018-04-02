// Framework
import React, {Component} from 'react';
import { connect } from "react-redux";
import Highcharts from 'highcharts';

// Constants
import {CHARTTYPE_PROGRESS} from "./../constants/ChartTypes";

const mapStateToProps = state => {
    return {
        chart: state.chart,
        chartButton: state.chartButton
    };
};

class ChartComp extends Component {
    componentDidUpdate () {
        let chartType = this.props.chartButton[this.props.chartButton.length - 1];
        switch (chartType) {
            case CHARTTYPE_PROGRESS:
                let last = this.props.chart[this.props.chart.length - 1];
                if (last) {
                    Highcharts.chart('chart', {
                        title: {
                            text: last.project_name
                        },
                        subtitle: {
                            text: 'Progress'
                        },
                        xAxis: {
                            title: {
                                text: "Date"
                            },
                            categories: last.categories
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
                        series: last.series
                    });
                }
                break;
        }
    }

    render() {
        let chartType = this.props.chartButton[this.props.chartButton.length - 1];
        switch (chartType) {
            case CHARTTYPE_PROGRESS:
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
