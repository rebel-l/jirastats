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

    render(){
        this.renderCounter++;
        if(this.renderCounter == 1){
            // we need to wait with rendering until first update is received
            return null;
        }

        let id = "PieChartTable-" +  this.projectId + "-" + this.actualChartType;
        let data = this.getData();
        let pieChartData = {
            title: "A project",
            subtitle: data.name,
            seriesData: data.data_chart
        };
        let tableData = {
            name: data.name,
            header: [data.name, "Count"],
            rows: data.data_table
        };
        return (
            <div key={id} id={id}>
                <ChartPie key={id + "-Chart-" + data.name} type={this.actualChartType} data={pieChartData}/>
                <Table key={id + "-Table-" + data.name} type={this.actualChartType} data={tableData}/>
            </div>
        );
    }
}

const PieChartTable = connect(mapStateToProps)(PieChartTableComp);
export default PieChartTable;
