import {CHART} from "../constants/actionTypes";

const Chart = chart => ({
    type: CHART,
    payload: chart
});

export default Chart;