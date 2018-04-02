// Framework
import React, {Component} from "react";

// Components
import ProjectSelect from "./ProjectSelect";
import Chart from "./Chart";
import ChartButton from "./ChartButton";

class App extends Component {
    render() {
        return (
            <div>
                <ProjectSelect/>
                <div>
                    <ChartButton name="Progress" chartType="progress"/>
                    <ChartButton name="Speed" chartType="speed"/>
                    <ChartButton name="Forecast" chartType="forecast"/>
                    <ChartButton name="Open Tasks" chartType="opentask"/>
                </div>
                <Chart/>
            </div>
        );
    }
}

export default App;