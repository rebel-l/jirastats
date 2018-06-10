// Framework
import React, {Component} from 'react';
import { connect } from "react-redux";
import Highcharts from "highcharts";

// Style
import style from "./Forecast.scss";

// Components
import Table from "./../Table";
import Summary from "./../Summary";

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

    getSpeed(){
        return {
            name: "Speed",
            header: ["Kind of Speed", "Tickets Per Day", "Tickets Per Week"],
            rows: this.getData().speed
        };
    }

    getSummary(){
        let summary = this.getData().summary;
        if (summary.days === -1 || summary.weeks === -1) {
            summary.days = "infinite";
            summary.weeks = "infinite";
            summary.last_day = "n/a";
            summary.last_week = "n/a";
        }
        return summary;
    }

    render(){
        this.renderCounter++;
        if(this.renderCounter === 1){
            // we need to wait with rendering until first update is received
            return null;
        }

        let options = {
            formatFloat: 2
        };

        return (
            <div key={this.id} id={this.id} className={style.forecast}>
                <div id={this.chartId} className={style.chart}/>
                <div>
                    <Table key={this.id + "-Table"} type={this.actualChartType} data={this.getSpeed()} options={options}/>
                </div>
                <Summary data={this.getSummary()}/>
            </div>
        );
    }
}

const Forecast = connect(mapStateToProps)(ForecastComp);
export default Forecast;
