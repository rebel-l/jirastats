// Constants
import {COMPARISON_SELECT_ACTION} from "../constants/ActionTypes";

let initialState = [{
    left: "",
    right: ""
}];

const Reducer = (state = initialState, action) => {
    switch (action.type) {
        case COMPARISON_SELECT_ACTION:
            let newState = Object.assign({}, state[state.length - 1]);
            newState[action.payload.version] = action.payload.value;
            return [...state, newState];
        default:
            return state;
    }
};

export default Reducer;