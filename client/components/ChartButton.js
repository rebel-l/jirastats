// Framework
import React, {Component} from 'react';
import { connect } from "react-redux";

// Style
import style from "./ChartButton.scss";

// Actions
import ChartButtonAction from "../actions/ChartButton";

const mapDispatchToProps = dispatch => {
    return {
        buttonClick: chartButton => dispatch(ChartButtonAction(chartButton))
    };
};

const mapStateToProps = state => {
    return { buttonClicked: state.chartButton };
};

class ChartButtonComp extends Component {
    constructor (props) {
        super(props);
        this.name = props.name;
        this.chartType = props.chartType;

        if(props.first === true){
            this.isFirst = " " + style.isFirst;
        } else {
            this.isFirst = "";
        }

        this.handleClick = this.handleClick.bind(this)
    }

    handleClick() {
        this.props.buttonClick(this.chartType);
    }

    render() {
        let last = this.props.buttonClicked[this.props.buttonClicked.length - 1];
        let isActive = "";
        if (last === this.chartType) {
            isActive = " " + style.isActive;
        }

        return (
            <button onClick={this.handleClick} className={style.chartButton + this.isFirst + isActive}>{this.name}</button>
        )
    }
}

const ChartButton = connect(mapStateToProps, mapDispatchToProps)(ChartButtonComp);

export default ChartButton