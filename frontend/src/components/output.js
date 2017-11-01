import React, { Component } from 'react';
import { Panel } from 'react-bootstrap';

class Output extends Component {
    constructor(props) {
        super();
        this.state = {
            build: {}
        };
    }

    render() {

        return (
            <Panel header={this.props.output.name + " on " + this.props.output.file}>
                { this.props.output.step }
                <pre>
                    {this.props.output.output}
                </pre>
            </Panel>
        );
    }
}

export default Output;
