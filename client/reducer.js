// Framework
import { combineReducers } from 'redux';

// Components
import ChartButton from "./reducers/ChartButton";
import ChartContainer from "./reducers/ChartContainer";
import Project from "./reducers/Project";

const Reducer = combineReducers({
    chartButton: ChartButton,
    chartContainer: ChartContainer,
    project: Project
});

export default Reducer;