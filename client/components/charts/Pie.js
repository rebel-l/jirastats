// Framework
import React, {Component} from 'react';
import Highcharts from 'highcharts';

class ChartPie extends Component {
    constructor(props){
        super(props);
        this.data = props.data;
        this.chartType = props.type;
        this.id = "PieChart-" + this.chartType + "-" + this.data.subtitle;
    }

    componentDidMount(){
        let options = {
            chart: {
                plotBackgroundColor: null,
                plotBorderWidth: null,
                plotShadow: false,
                type: 'pie'
            },
            title: {
                text: this.data.title
            },
            subtitle: {
                text: this.data.subtitle
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
            series: [{
                name: 'Tickets',
                colorByPoint: true,
                data: this.data.seriesData
            }]
        };
        Highcharts.chart(this.id, options);
    }

    render(){
        return (
            <div key={this.id} id={this.id} />
        );
    }
}

export default ChartPie