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

    handleSelect(){
        console.log("selected something")
    }

    render() {
        return (
            <div>
                <label>Select Project: </label>
                <select onChange={this.handleSelect}>
                    <option key={0} id={0}></option>
                    {this.state.projects.map(project =>
                        <option key={project.id} id={project.id}>{project.name}</option>
                    )}
                </select>
            </div>
        );
    }
}

export default ProjectSelect