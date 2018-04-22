// Framework
import React, {Component} from 'react';
import { connect } from "react-redux";
import Highcharts from 'highcharts';

const mapStateToProps = state => {
    return {
        newChartType: state.chartButton,
        data: state.chartContainer
    };
};

class ChartLineComp extends Component {
    constructor(props){
        super(props);
        this.actualChartType = props.type;
        this.options = props.options;
    }

    shouldComponentUpdate(nextProps, nextStates){
        this.id = "LineChart-" + nextProps.project + "-" + this.actualChartType;
        return true;
    }

    componentDidUpdate(){
        let data = this.getData();
        let options = {
            chart: {
                type: 'line'
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
        Highcharts.chart(this.id, options);
    }

    getData(){
        return this.props.data[this.props.data.length -1];
    };

    render(){
        let newChartType = this.props.newChartType[this.props.newChartType.length - 1];
        if (newChartType !== this.actualChartType) {
            // we get the data for a different chart
            return null;
        }

        return (
            <div key={this.id} id={this.id} />
        );
    }
}

const ChartLine = connect(mapStateToProps)(ChartLineComp);
export default ChartLine