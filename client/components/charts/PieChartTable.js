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

    renderChild(id, data) {
        return (
            <div key={id + "-" + data.name} id={id + "-" + data.name}>
                <ChartPie key={id + "-Chart-" + data.name} type={this.actualChartType} data={this.getPieChartData(data)}/>
                <Table key={id + "-Table-" + data.name} type={this.actualChartType} data={this.getTableData(data)}/>
            </div>
        );
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
                {this.renderChild(id, data.priorities)}
                {this.renderChild(id, data.status)}
                {this.renderChild(id, data.issue_types)}
                {this.renderChild(id, data.tech_debt)}
                {this.renderChild(id, data.components)}
                {this.renderChild(id, data.labels)}
            </div>
        );
    }
}

const PieChartTable = connect(mapStateToProps)(PieChartTableComp);
export default PieChartTable;
