// Framework
import React, {Component} from 'react';
import { connect } from "react-redux";
import Highcharts from "highcharts";

const mapStateToProps = state => {
    return {
        data: state.chartContainer
    };
};

class ForecastComp extends Component {
    constructor(props){
        super(props);
        this.actualChartType = props.type;
        this.projectId = props.project;
        this.renderCounter = 0;
        this.id = "Forecast" +  this.projectId + "-" + this.actualChartType;
        this.chartId = this.id + "-chart"
    }

    getData(){
        return this.props.data[this.props.data.length -1];
    };

    componentDidUpdate(){
        let data = this.getData().chart;
        let options = {
            chart: {
                type: "line"
            },
            title: {
                text: data.project_name
            },
            subtitle: {
                text: "Forecast"
            },
            xAxis: {
                title: {
                    text: "Date"
                },
                categories: data.categories
            },
            yAxis: {
                title: {
                    text: "Number of open tickets"
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
        Highcharts.chart(this.chartId, options);
    }

    render(){
        this.renderCounter++;
        if(this.renderCounter === 1){
            // we need to wait with rendering until first update is received
            return null;
        }

        return (
            <div key={this.id} id={this.id}>
                <div id={this.chartId}/>
            </div>
        );
    }
}

const Forecast = connect(mapStateToProps)(ForecastComp);
export default Forecast;
