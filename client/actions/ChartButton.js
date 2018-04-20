import {CHART_TYPE_BUTTON_ACTION} from "../constants/ActionTypes";

const ChartButton = chartButton => ({
    type: CHART_TYPE_BUTTON_ACTION,
    payload: chartButton
});

export default ChartButton;