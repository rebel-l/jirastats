import {CHART_BUTTON} from "../constants/ActionTypes";

const Reducer = (state = ["progress"], action) => {
    switch (action.type) {
        case CHART_BUTTON:
            return [ ...state, action.payload ];
        default:
            return state;
    }
};

export default Reducer;