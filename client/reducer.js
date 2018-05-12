// Framework
import { combineReducers } from 'redux';

// Components
import ChartButton from "./reducers/ChartButton";
import ChartContainer from "./reducers/ChartContainer";
import ComparisonSelector from "./reducers/ComparisonSelector";
import Project from "./reducers/Project";

const Reducer = combineReducers({
    chartButton: ChartButton,
    chartContainer: ChartContainer,
    comparisonSelector: ComparisonSelector,
    project: Project
});

export default Reducer;