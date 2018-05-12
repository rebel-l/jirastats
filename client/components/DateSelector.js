// Framework
import React, {Component} from 'react';

class DateSelectorComp extends Component {
    constructor(props){
        super(props);
        this.id = props.id + "-select";
        this.data = props.data;
        this.version = props.version;
    }

    getOption(date){
        return (
            <option key={this.id + "-" + date.value} value={date.value}>{date.name}</option>
        );
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
            <select key={this.id} id={this.id}>
                {children}
            </select>
        );
    }
}

export default DateSelectorComp;