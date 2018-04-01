// Framework
import axios from 'axios';
import React, {Component} from 'react';
import { connect } from "react-redux";

// Actions
import Chart from "../actions/Chart";

const mapDispatchToProps = dispatch => {
    return {
        projectSelect: chart => dispatch(Chart(chart))
    };
};

class Project extends Component {
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
        if (event.target.value !== 0) {
            axios.get(`/data/stats/${event.target.value}`).then(res => {
                this.props.projectSelect(res.data);
            })
        }
    }

    render() {
        return (
            <div>
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

const ProjectSelect = connect(null, mapDispatchToProps)(Project);

export default ProjectSelect