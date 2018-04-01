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
            // TODO: get the stats data for project
            // TODO: send the data to redux
            // TODO: let redux charts listen to redux event and send chart options to chart component
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