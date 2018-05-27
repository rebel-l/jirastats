// Framework
import React, {Component} from "react";

// Components
import ChartContainer from "./ChartContainer";
import ChartButton from "./ChartButton";
import Footer from "./Footer";
import ProjectSelect from "./ProjectSelect";

// Constants
import {CHARTTYPE_PROGRESS} from "./../constants/ChartTypes";
import {CHARTTYPE_SPEED} from "./../constants/ChartTypes";
import {CHARTTYPE_FORECAST} from "./../constants/ChartTypes";
import {CHARTTYPE_OPENTICKETS} from "./../constants/ChartTypes";
import {CHARTTYPE_COMPARISON} from "./../constants/ChartTypes";

class App extends Component {
    render() {
        let appId = 'Application';
        let chartTypeSelectorId = 'ChartTypeSelector';
        return (
            <div key={appId} id={appId}>
                <ProjectSelect/>
                <div key={chartTypeSelectorId} id={chartTypeSelectorId}>
                    <ChartButton name="Progress" chartType={CHARTTYPE_PROGRESS}/>
                    <ChartButton name="Speed" chartType={CHARTTYPE_SPEED}/>
                    <ChartButton name="Forecast" chartType={CHARTTYPE_FORECAST}/>
                    <ChartButton name="Open Tickets" chartType={CHARTTYPE_OPENTICKETS}/>
                    <ChartButton name="Comparison" chartType={CHARTTYPE_COMPARISON}/>
                </div>
                <ChartContainer/>
                <Footer/>
            </div>
        );
    }
}

export default App;