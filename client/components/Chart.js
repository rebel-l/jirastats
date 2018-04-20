// Framework
import axios from "axios/index";
import React, {Component} from 'react';
import { connect } from "react-redux";
import Highcharts from 'highcharts';

// Constants
import {CHARTTYPE_PROGRESS} from "./../constants/ChartTypes";
import {CHARTTYPE_SPEED} from "./../constants/ChartTypes";
import {CHARTTYPE_OPENTICKETS} from "./../constants/ChartTypes";

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

        // axios.get(`/data/stats/progress/${project}`).then(res => { // TODO: temporary static test
        axios.get(`/data/stats/${chartType}/${project}`).then(res => {
            switch (chartType) {
                case CHARTTYPE_PROGRESS:
                    Highcharts.chart('chart', this.getChartOptions("Progress", "line", res.data, "Date"));
                    break;
                case CHARTTYPE_SPEED:
                    Highcharts.chart('chart', this.getChartOptions("Speed", "column", res.data, "Week"));
                    break;
                case CHARTTYPE_OPENTICKETS:
                    let chartData = {
                        project_name: "A project",
                        series: [{
                            name: 'Tickets',
                            colorByPoint: true,
                            data: res.data.data_chart /*[{
                                name: 'Catalog Service',
                                y: 42.5,
                                sliced: true,
                                // selected: true
                            }, {
                                name: 'Order Service',
                                y: 14.5
                            }, {
                                name: 'Wishlist Service',
                                y: 1
                            }, {
                                name: 'Customer Service',
                                y: 11
                            }, {
                                name: 'Catalog Importer',
                                y: 31
                            }]*/
                        }]
                    };
                    Highcharts.chart('pieChart', this.getPieChartOptions(res.data.name, chartData));
                    break;
            }
        });
    }

    getChartOptions(subtitle, type, chartData, xAxisTitle) {
        return {
            chart: {
                type: type
            },
            title: {
                text: chartData.project_name
            },
            subtitle: {
                text: subtitle
            },
            xAxis: {
                title: {
                    text: xAxisTitle
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

    getPieChartOptions(subtitle, chartData) {
        return {
            chart: {
                plotBackgroundColor: null,
                plotBorderWidth: null,
                plotShadow: false,
                type: 'pie'
            },
            title: {
                text: chartData.project_name
            },
            subtitle: {
                text: subtitle
            },
            tooltip: {
                pointFormat: '{series.name}: <b>{point.percentage:.1f}%</b>'
            },
            plotOptions: {
                pie: {
                    allowPointSelect: true,
                    cursor: 'pointer',
                    dataLabels: {
                        enabled: true,
                        format: '<b>{point.name}</b>: {point.percentage:.1f} %',
                        style: {
                            color: (Highcharts.theme && Highcharts.theme.contrastTextColor) || 'black'
                        }
                    }
                }
            },
            series: chartData.series
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
            case CHARTTYPE_OPENTICKETS:
                // TODO: add data table and have multiple charts
                return (
                    <div key={"pieChartContainer"}>
                        <div id="pieChart" />
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
