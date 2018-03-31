// Framework
import React, {Component} from "react";

// Components
import ProjectSelect from "./ProjectSelect";
import Chart from "./Chart";

class App extends Component {
    render() {
        return (
            <div>
                <ProjectSelect/>
                <Chart/>
            </div>
    );
    }
}

export default App;