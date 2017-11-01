import React, { Component } from 'react';
import { Nav, NavItem, } from 'react-bootstrap';

class Builds extends Component {
    constructor() {
        super();
        this.state = {
            builds: []
        };
    }

    componentDidMount() {
        fetch(`https://continuous-ino.eyskens.me/api/build/all`)
        .then(result=>result.json())
        .then(builds=>this.setState({builds}))
    }

    handleSelect(selectedKey) {
        alert(`selected ${selectedKey}`);
    }
    render() {
        var buildsList = this.state.builds.map((build, key) => {
            return <NavItem key={ key }>{build.id} - {build.repo}: {build.sha}</NavItem>;
          })


        return (
            <div>
                <h2>Builds</h2>
                <Nav bsStyle="pills" stacked activeKey={1} onSelect={this.handleSelect}>
                    { buildsList }
                </Nav>
            </div>
        );
    }
}

export default Builds;
