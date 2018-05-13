// Framework
import React, {Component} from 'react';
import { connect } from "react-redux";

// Actions
import ComparisonSelector from "../actions/ComparisonSelector";

const mapDispatchToProps = dispatch => {
    return {
        comparisonSelector: comparisonSelector => dispatch(ComparisonSelector(comparisonSelector))
    };
};

class DateSelectorComp extends Component {
    constructor(props){
        super(props);
        this.id = props.id + "-select";
        this.data = props.data;
        this.version = props.version;
        this.handleSelect = this.handleSelect.bind(this);
    }

    getOption(date){
        return (
            <option key={this.id + "-" + date.value} value={date.value}>{date.name}</option>
        );
    }

    handleSelect(event){
        if(event.target.value !== ""){
            this.props.comparisonSelector({
               version: this.version,
               value: event.target.value
            });
        }
    }

    render(){
        let children = [];
        children.push(this.getOption({
            name: `... select date ${this.version} ...`,
            value: ""
        }));
        for(let i = 0; i < this.data.length; i++){
            children.push(this.getOption(this.data[i]));
        }

        return(
            <select key={this.id} id={this.id} onChange={this.handleSelect}>
                {children}
            </select>
        );
    }
}

const DateSelector = connect(null, mapDispatchToProps)(DateSelectorComp);

export default DateSelector;