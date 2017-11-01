import React, { Component } from 'react';
import Output from './output'

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
        if (this.state.build.output != null) {
            var outputList = this.state.build.output.map((output, key) => {
                return (
                    <Output key={key} output={output} />
                );
            })
        } else {
            var outputList = <p>Waiting for output...</p>
        }

        return (
            <div>
                <h2>Build #{this.props.match.params.id}</h2>
                <h3>{this.state.build.repo} - {this.state.build.sha}</h3>
                { outputList }
            </div>
        );
    }
}

export default Build;
