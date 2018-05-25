// Framework
import React, {Component} from 'react';

class Table extends Component {
    constructor(props){
        super(props);
        this.data = props.data;
        this.chartType = props.type;
        this.id = "Table-" + this.chartType + "-" + this.data.name;
    }

    render(){
        let header = [];
        for (let i = 0; i < this.data.header.length; i++){
            header.push(<td key={this.id + '-header-' + i}>{this.data.header[i]}</td>)
        }

        let rows = [];
        if (this.data.rows !== undefined && this.data.rows !== null){
            for(let i = 0; i < this.data.rows.length; i++){
                let cells = [];
                let j = 0;
                for(let key in this.data.rows[i]){
                    let cellId = `${this.id}-cell-${i}-${j}`;
                    if(this.data.rows[i].hasOwnProperty(key)) {
                        cells.push(<td key={cellId}>{this.data.rows[i][key]}</td>);
                    }
                    j++;
                }
                rows.push(<tr key={this.id + '-row-' + i}>{cells}</tr>);
            }
        }

        return (
            <table key={this.id}>
                <thead key={this.id + '-thead'}>
                <tr key={this.id + '-header'}>
                    {header}
                </tr>
                </thead>
                <tbody key={this.id + '-tbody'}>
                {rows}
                </tbody>
            </table>
        )
    }
}

export default Table