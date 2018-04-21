// Framework
import React, {Component} from 'react';

class Table extends Component {
    constructor(props){
        super(props);
        this.header = ['Status', 'Count'];
        this.data = [{
            name: 'Open',
            value: 2
        }];

        this.data = [
          ['Open', 2],
          ['In Progress', 1]
        ];
        this.header = props.options.header;
        this.data = props.options.data;
    }

    render(){
        let header = [];
        for (let i = 0; i < this.header.length; i++){
            header.push(<td key={'header-' + i}>{this.header[i]}</td>)
        }

        let rows = [];
        for(let i = 0; i < this.data.length; i++){
            let cells = [];
            let j = 0;
            for(let key in this.data[i]){
                let cellId = `cell-${i}-${j}`;
                if(this.data[i].hasOwnProperty(key)) {
                    cells.push(<td key={cellId}>{this.data[i][key]}</td>);
                }
                j++;
            }
            rows.push(<tr key={'row-' + i}>{cells}</tr>);
        }

        return (
            <table key={'table'}>
                <thead key={'thead'}>
                    <tr key={'header'}>
                        {header}
                    </tr>
                </thead>
                <tbody key={'tbody'}>
                    {rows}
                </tbody>
            </table>
        )
    }
}

export default Table