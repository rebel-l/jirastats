// Framework
import axios from 'axios';
import React, {Component} from 'react';
import { connect } from "react-redux";

// Actions
import Project from "../actions/Project"

const mapDispatchToProps = dispatch => {
    return {
        project: project => dispatch(Project(project))
    };
};

class ProjectSelectComponent extends Component {
    constructor(props) {
        super(props);

        this.state = {
            projects: []
        };

        this.handleSelect = this.handleSelect.bind(this);
    }

    componentDidMount() {
        axios.get("/projects").then(res => {
            this.setState(res.data);
        });
    }

    handleSelect(event){
        if (event.target.value !== "0") {
            this.props.project(event.target.value);
        }
    }

    render() {
        let id = 'ProjectSelector';
        return (
            <div key={id} id={id}>
                <label>Select Project: </label>
                <select onChange={this.handleSelect} value={this.state.value}>
                    <option key={0} value={0}>... select project ...</option>
                    {this.state.projects.map(project =>
                        <option key={project.id} value={project.id}>{project.name}</option>
                    )}
                </select>
            </div>
        );
    }
}

const ProjectSelect = connect(null, mapDispatchToProps)(ProjectSelectComponent);

export default ProjectSelect