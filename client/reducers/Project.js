import {PROJECT_SELECT_ACTION} from "../constants/ActionTypes";

const Reducer = (state = [], action) => {
    switch (action.type) {
        case PROJECT_SELECT_ACTION:
            return [ ...state, action.payload ];
        default:
            return state;
    }
};

export default Reducer;