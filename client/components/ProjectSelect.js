// Framework
import axios from 'axios';
import React, {Component} from 'react';

class ProjectSelect extends Component {
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
        if (event.target.value != 0) {
            console.log("selected: " + event.target.value)
        }
    }

    render() {
        return (
            <div>
                <label>Select Project: </label>
                <select onChange={this.handleSelect} value={this.state.value}>
                    <option key={0} value={0}></option>
                    {this.state.projects.map(project =>
                        <option key={project.id} value={project.id}>{project.name}</option>
                    )}
                </select>
            </div>
        );
    }
}

export default ProjectSelect