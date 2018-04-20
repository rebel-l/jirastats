// Framework
import React, {Component} from "react";

// Components
import ProjectSelect from "./ProjectSelect";
import Chart from "./Chart";
import ChartButton from "./ChartButton";

// Constants
import {CHARTTYPE_PROGRESS} from "./../constants/ChartTypes";
import {CHARTTYPE_SPEED} from "./../constants/ChartTypes";
import {CHARTTYPE_FORECAST} from "./../constants/ChartTypes";
import {CHARTTYPE_OPENTICKETS} from "./../constants/ChartTypes";
import {CHARTTYPE_COMPARISON} from "./../constants/ChartTypes";

class App extends Component {
    render() {
        return (
            <div>
                <ProjectSelect/>
                <div>
                    <ChartButton name="Progress" chartType={CHARTTYPE_PROGRESS}/>
                    <ChartButton name="Speed" chartType={CHARTTYPE_SPEED}/>
                    <ChartButton name="Forecast" chartType={CHARTTYPE_FORECAST}/>
                    <ChartButton name="Open Tickets" chartType={CHARTTYPE_OPENTICKETS}/>
                    <ChartButton name="Comparison" chartType={CHARTTYPE_COMPARISON}/>
                </div>
                <Chart/>
            </div>
        );
    }
}

export default App;