// Framework
import React, {Component} from 'react';
import Highcharts from 'highcharts';

class Chart extends Component {
    componentDidMount() {
        console.log("show chart");
        Highcharts.chart('chart', {
            title: {
                text: 'Project Name'
            },
            subtitle: {
                text: 'Progress'
            },
            xAxis: {
                title: {
                    text: "Days"
                },
                categories: ["26.03.2018", "27.03.2018", "28.03.2018", "29.03.2018"]
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
            series: [{
                name: 'Open',
                data: [100, 92, 83, 88]
            }, {
                name: 'Closed',
                data: [0, 8, 15, 4]
            }, {
                name: 'New',
                data: [0, 0, 6, 9]
            }]
        });
    }

    render() {
        return (
            <div id="chart" />
        )
    }
}

export default Chart
