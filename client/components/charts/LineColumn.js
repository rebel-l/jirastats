// Framework
import React, {Component} from 'react';
import { connect } from "react-redux";
import Highcharts from 'highcharts';

const mapStateToProps = state => {
    return {
        data: state.chartContainer
    };
};

class ChartLineComp extends Component {
    constructor(props){
        super(props);
        this.options = props.options;
    }

    componentDidUpdate(){
        let data = this.getData();
        let options = {
            chart: {
                type: this.options.type
            },
            title: {
                text: data.project_name
            },
            subtitle: {
                text: this.options.title
            },
            xAxis: {
                title: {
                    text: this.options.xAxis
                },
                categories: data.categories
            },
            yAxis: {
                title: {
                    text: this.options.yAxis
                }
            },
            legend: {
                layout: 'vertical',
                align: 'right',
                verticalAlign: 'middle'
            },
            series: data.series,
            credits: false
        };
        Highcharts.chart(this.options.id, options);
    }

    getData(){
        return this.props.data[this.props.data.length -1];
    };

    render(){
        return (
            <div key={this.options.id} id={this.options.id} />
        );
    }
}

const ChartLine = connect(mapStateToProps)(ChartLineComp);
export default ChartLine