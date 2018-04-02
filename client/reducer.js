// Framework
import { combineReducers } from 'redux';

// Components
import Chart from './reducers/Chart';
import ChartButton from "./reducers/ChartButton";

const Reducer = combineReducers({
    chart: Chart,
    chartButton: ChartButton
});

export default Reducer;