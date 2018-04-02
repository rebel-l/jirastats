import {CHART_BUTTON} from "../constants/actionTypes";

const ChartButton = chartButton => ({
    type: CHART_BUTTON,
    payload: chartButton
});

export default ChartButton;