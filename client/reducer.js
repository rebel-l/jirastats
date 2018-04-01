import { combineReducers } from 'redux';
import Chart from './reducers/Chart';

const Reducer = combineReducers({
    chart: Chart
});

export default Reducer;