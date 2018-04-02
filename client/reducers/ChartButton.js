// Constants
import {CHART_BUTTON} from "../constants/ActionTypes";
import {CHARTTYPE_PROGRESS} from "./../constants/ChartTypes";

const Reducer = (state = [CHARTTYPE_PROGRESS], action) => {
    switch (action.type) {
        case CHART_BUTTON:
            return [ ...state, action.payload ];
        default:
            return state;
    }
};

export default Reducer;