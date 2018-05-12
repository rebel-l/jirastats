import {COMPARISON_SELECT_ACTION} from "../constants/ActionTypes";

const ComparisonSelector = comparisonSelector => ({
    type: COMPARISON_SELECT_ACTION,
    payload: comparisonSelector
});

export default ComparisonSelector;