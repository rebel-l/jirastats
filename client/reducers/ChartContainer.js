// Constants
import {CHART_CONTAINER_DATA_ACTION} from "../constants/ActionTypes";

const Reducer = (state = [], action) => {
    switch (action.type) {
        case CHART_CONTAINER_DATA_ACTION:
            return [ ...state, action.payload ];
        default:
            return state;
    }
};

export default Reducer;