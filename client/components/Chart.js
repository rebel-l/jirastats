// Framework
import React, {Component} from 'react';
import { connect } from "react-redux";
import Highcharts from 'highcharts';

const mapStateToProps = state => {
    return { chart: state.chart };
};

class ChartComp extends Component {
    componentDidUpdate () {
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
    }

    render() {
        return (
            <div id="chart" />
        )
    }
}

const Chart = connect(mapStateToProps)(ChartComp);

export default Chart
