import {CHART} from "../constants/ActionTypes";

const Reducer = (state = [], action) => {
    switch (action.type) {
        case CHART:
            return [ ...state, action.payload ];
        default:
            return state;
    }
};


export default Reducer;