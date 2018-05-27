// Framework
import React, {Component} from "react";

// Style
import style from "./App.scss";

// Components
import ChartContainer from "./ChartContainer";
import ChartButton from "./ChartButton";
import Header from "./Header";
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
            <div key={appId} id={appId} className={style.spacer}>
                <Header/>
                <nav>
                    <ProjectSelect/>
                    <div key={chartTypeSelectorId} id={chartTypeSelectorId} className={style.chartSelector}>
                        <label>Select Chart:</label>
                        <ChartButton name="Progress" chartType={CHARTTYPE_PROGRESS} first={true}/>
                        <ChartButton name="Speed" chartType={CHARTTYPE_SPEED}/>
                        <ChartButton name="Forecast" chartType={CHARTTYPE_FORECAST}/>
                        <ChartButton name="Open Tickets" chartType={CHARTTYPE_OPENTICKETS}/>
                        <ChartButton name="Comparison" chartType={CHARTTYPE_COMPARISON}/>
                    </div>
                </nav>
                <ChartContainer/>
                <Footer/>
            </div>
        );
    }
}

export default App;