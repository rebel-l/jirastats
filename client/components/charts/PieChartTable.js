// Framework
import React, {Component} from 'react';
import { connect } from "react-redux";

// Components
import ChartPie from "./Pie";
import Table from "./../Table";

const mapStateToProps = state => {
    return {
        data: state.chartContainer
    };
};

class PieChartTableComp extends Component {
    constructor(props){
        super(props);
        this.actualChartType = props.type;
        this.projectId = props.project;
        this.renderCounter = 0;
    }

    getData(){
        return this.props.data[this.props.data.length -1];
    };

    getPieChartData(data){
        let pieChartData = {
            title: "A project",
            subtitle: data.name,
            seriesData: data.data_chart
        };
        return pieChartData;
    }

    getTableData(data){
        let tableData = {
            name: data.name,
            header: [data.name, "Count"],
            rows: data.data_table
        };
        return tableData;
    }

    render(){
        this.renderCounter++;
        if(this.renderCounter == 1){
            // we need to wait with rendering until first update is received
            return null;
        }

        let id = "PieChartTable-" +  this.projectId + "-" + this.actualChartType;
        let data = this.getData();

        return (
            <div key={id} id={id}>
                <div key={id + "-" + data.priority.name} id={id + "-" + data.priority.name}>
                    <ChartPie key={id + "-Chart-" + data.priority.name} type={this.actualChartType} data={this.getPieChartData(data.priority)}/>
                    <Table key={id + "-Table-" + data.priority.name} type={this.actualChartType} data={this.getTableData(data.priority)}/>
                </div>
                <div key={id + "-" + data.status.name} id={id + "-" + data.status.name}>
                    <ChartPie key={id + "-Chart-" + data.status.name} type={this.actualChartType} data={this.getPieChartData(data.status)}/>
                    <Table key={id + "-Table-" + data.status.name} type={this.actualChartType} data={this.getTableData(data.status)}/>
                </div>
            </div>
        );
    }
}

const PieChartTable = connect(mapStateToProps)(PieChartTableComp);
export default PieChartTable;
