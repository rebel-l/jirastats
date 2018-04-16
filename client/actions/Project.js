import {PROJECT_SELECT_ACTION} from "../constants/actionTypes";

const Project = project => ({
    type: PROJECT_SELECT_ACTION,
    payload: project
});

export default Project;