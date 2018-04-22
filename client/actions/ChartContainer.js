import {CHART_CONTAINER_DATA_ACTION} from "../constants/ActionTypes";

const ChartContainer = chartContainer => ({
    type: CHART_CONTAINER_DATA_ACTION,
    payload: chartContainer
});

export default ChartContainer;