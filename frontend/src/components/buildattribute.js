import React, { Component } from 'react';
import FontAwesome from 'react-fontawesome';

class BuildAttribute extends Component {

    render() {
        if (this.props.build.running) {
            return <FontAwesome name='refresh' spin/>
        }

        if (this.props.build.errors === null || this.build.errors.length === 0) {
            return <FontAwesome name='check' style={{ color: 'green' }}/>
        }

        return <FontAwesome name='times' style={{ color: 'red' }}/>

    }
}

export default BuildAttribute;
