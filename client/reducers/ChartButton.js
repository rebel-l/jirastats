// Constants
import {CHART_TYPE_BUTTON_ACTION} from "../constants/ActionTypes";
import {CHARTTYPE_PROGRESS} from "./../constants/ChartTypes";

const Reducer = (state = [CHARTTYPE_PROGRESS], action) => {
    switch (action.type) {
        case CHART_TYPE_BUTTON_ACTION:
            return [ ...state, action.payload ];
        default:
            return state;
    }
};

export default Reducer;