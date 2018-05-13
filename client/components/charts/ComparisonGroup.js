// Framework
import React, {Component} from 'react';
import { connect } from "react-redux";
import axios from "axios/index";

// Components
import ChartPie from "./Pie";
import Table from "./../Table";

const mapStateToProps = state => {
    return {
        dates: state.comparisonSelector
    };
};

class ComparisonGroupComp extends Component{
    constructor(props){
        super(props);
        this.dataRecieved = false;
        this.id = props.id;
        this.state = {
            projectId: props.projectId,
            data: null
        };
    }

    shouldComponentUpdate(nextProps){
        let lastDates = nextProps.dates[nextProps.dates.length - 1];
        return !(lastDates.left === "" || lastDates.right === "")
    }

    componentDidUpdate(){
        if(this.dataRecieved === true){
            this.dataRecieved = false;
            return;
        }

        this.doRequest();
    }

    doRequest(){
        let dates = this.getComparisonSelectorDates();
        axios.get(`/data/stats/comparison/${this.state.projectId}/${dates.left}/${dates.right}`).then(res => {
            this.dataRecieved = true;
            let newState = Object.assign({}, this.state);
            newState.data = res.data;
            this.setState(newState);
        });
    }

    getComparisonSelectorDates(){
        return this.props.dates[this.props.dates.length - 1];
    }

    static getPieData(title, subtitle, series){
        return {
            title: title,
            subtitle: subtitle,
            seriesData: series
        }
    }

    static getTableData(versions, data){
        return {
            header: ["Name", versions.left, versions.right, "Diff"],
            rows: data
        }
    }

    static getChild(id, title, versions, stats){
        return (
            <div key={id}>
                <ChartPie type={versions.left} data={ComparisonGroupComp.getPieData(title + " - " + versions.left, stats.name, stats.left)}/>
                <ChartPie type={versions.right} data={ComparisonGroupComp.getPieData(title + " - " + versions.right, stats.name, stats.right)}/>
                <Table type={"comparison"} data={ComparisonGroupComp.getTableData(versions, stats.diff)}/>
            </div>
        );
    }

    getId(){
        let versions = this.getComparisonSelectorDates();
        return this.id + "-" + versions.left + "-" + versions.right;
    }

    render(){
        if(this.state.data === null){
            return(
                <p>Select the dates for the two versions of project {this.state.projectId}.</p>
            );
        } else if(this.dataRecieved === false) {
            return(
                <p>Loading ...</p>
            );
        }

        let title = this.state.data.project.name;
        let children = [];
        for(let prop in this.state.data.stats) {
            if(!this.state.data.stats.hasOwnProperty(prop)) continue;
            let childId = this.id + "-" + prop;
            children.push(ComparisonGroupComp.getChild(childId, title, this.state.data.version, this.state.data.stats[prop]))
        }

        return(
          <div key={this.getId()} id={this.getId()}>
              {children}
          </div>
        );
    }
}

const ComparisonGroup = connect(mapStateToProps)(ComparisonGroupComp);
export default ComparisonGroup;