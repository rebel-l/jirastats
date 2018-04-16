// Framework
import { combineReducers } from 'redux';

// Components
import ChartButton from "./reducers/ChartButton";
import Project from "./reducers/Project";

const Reducer = combineReducers({
    chartButton: ChartButton,
    project: Project
});

export default Reducer;