import React, { Component } from 'react';
import Output from './output'
import BuildAttribute  from './buildattribute'

class Build extends Component {
    constructor() {
        super();
        this.state = {
            loading: true,
        }
    }
      
    componentDidMount() {
        fetch(`https://continuous-ino.eyskens.me/api/build/${this.props.match.params.id}`)
        .then(result=>result.json())
        .then(build=>this.setState({ build, loading: false }))
    }

    render() {
        if (this.state.loading) {
            return <p>Loading...</p>
        }

        var outputList
        if (this.state.build.output != null) {
            outputList = this.state.build.output.map((output, key) => {
                return (
                    <Output key={key} output={output} />
                );
            })
        } else {
            outputList = <p>Waiting for output...</p>
        }

        return (
            <div>
                <h2>Build #{this.props.match.params.id} <BuildAttribute build={this.state.build}/></h2>
                <h3>{this.state.build.repo} - {this.state.build.sha}</h3>
                { outputList }
            </div>
        );
    }
}

export default Build;
